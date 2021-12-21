package week07

import (
	"context"
	"fmt"
	"sync"
)

// Manager is responsible for receiving product orders
// and assigning them to employees. Manager is also responsible
// for receiving completed products, and listening for errors,
// from employees. Manager takes products that have been built
// by employees and returns them to the customer as a CompletedProduct.
type Manager struct {
	// non-exported fields (PRIVATE)
	// !YOU MAY NOT ACCESS THESE FIELDS IN YOUR TESTS!
	completed chan CompletedProduct
	errs      chan error
	jobs      chan *Product
	stopped   bool
	cancel    context.CancelFunc
	sync.RWMutex
	sync.Once
}

// NewManager will create a new Manager.
// This function should ALWAYS be used to
// create a new Manager.
func NewManager() *Manager {
	return &Manager{
		completed: make(chan CompletedProduct),
		errs:      make(chan error),
		jobs:      make(chan *Product),
	}
}

// Start will create new employees for the given count,
// and start listening for jobs and errors.
// Managers should be stopped using the Stop method
// when they are no longer needed.
func (m *Manager) Start(ctx context.Context, count int) (context.Context, error) {

	if count <= 0 {
		return nil, ErrInvalidEmployeeCount(count)
	}

	ctx, cancel := context.WithCancel(ctx)
	m.Lock()
	m.cancel = cancel
	m.Unlock()

	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println(ctx.Err())
		cancel()
		m.Stop()

	}(ctx)

	for i := 0; i < count; i++ {

		e := Employee(i + 1)
		go e.work(ctx, m)
	}

	//time.Sleep(time.Second * 10)
	return ctx, nil
}

// Assign will assign the given products to employees
// as employeess become available. An invalid product
// will return an error.
func (m *Manager) Assign(products ...*Product) error {
	m.RLock()
	if m.stopped {
		m.RUnlock()
		return ErrManagerStopped{}
	}
	m.RUnlock()

	// loop through each product and assign it to an employee
	for _, p := range products {
		// validate product
		if err := p.IsValid(); err != nil {
			return err
		}

		// assign product to employee
		// this will block until an employee becomes available
		m.Jobs() <- p
	}

	return nil
}

// Complete will wrap the employee and the product into
// a CompletedProduct. The will be passed down the Completed()
// channel as soon as a listener is available to receive it.
// Complete will error if the employee is invalid or
// if the product is not built.
func (m *Manager) Complete(e Employee, p *Product) error {
	// validate employee
	if err := e.IsValid(); err != nil {
		return err
	}

	// validate product is built
	if err := p.IsBuilt(); err != nil {
		return err
	}

	//p.Lock()
	cp := CompletedProduct{
		Employee: e,
		Product:  *p, // deference pointer to value type ype t
	}
	//p.Unlock()

	// Send completed product to Completed() channel
	// for a listener to receive it.
	// This will block until a listener is available.
	m.completedCh() <- cp

	return nil
}

// completedCh returns the channel for CompletedProducts
func (m *Manager) completedCh() chan CompletedProduct {
	m.Lock()
	defer m.Unlock()
	return m.completed
}

// Completed will return a channel that can be listened to
// for CompletedProducts.
// This is a read-only channel.
func (m *Manager) Completed() <-chan CompletedProduct {
	return m.completedCh()
}

// Jobs will return a channel that can be listened to
// for new products to be built.
func (m *Manager) Jobs() chan *Product {
	m.Lock()
	defer m.Unlock()
	return m.jobs
}

// Errors will return a channel that can be listened to
// and can be used to receive errors from employees.
func (m *Manager) Errors() chan error {
	m.Lock()
	defer m.Unlock()
	return m.errs
}

// Stop will stop the manager and clean up all resources.
func (m *Manager) Stop() {
	m.RLock()
	if m.stopped {
		m.RUnlock()
		return
	}
	m.RUnlock()

	m.Lock()
	m.stopped = true
	m.Unlock()

	// close all channels
	close(m.jobs)
	close(m.errs)

	//Might need this
	if m.completed != nil {
		close(m.completed)
	}
}

// snippet: example
func Run(ctx context.Context, count int, products ...*Product) ([]CompletedProduct, error) {
	// NOTE: this function should not be the one to create
	// the necessary contexts, time outs, signals, etc.
	// The Run method should not care about those concerns,
	// only its own.

	// TODO: implement this function
	// This function should run the manager with the given products
	// and return the results.
	m := NewManager()
	ctx, err := m.Start(ctx, count)
	if err != nil {
		return nil, err
	}

	//Go through the products and assign them to employees
	for i := range products {
		go m.Assign(products[i])
	}
	var completed []CompletedProduct
	go func() {

		//m.Lock()
		for cp := range m.Completed() {
			completed = append(completed, cp)

			if len(completed) >= len(products) {
				m.Stop()
			}
		}
		//m.Unlock()
	}()
	<-ctx.Done()
	m.RLock()
	defer m.RUnlock()
	return completed, nil
}
