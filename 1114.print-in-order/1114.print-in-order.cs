using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;

public class Foo {
    
    private readonly Semaphore first = new Semaphore(0, 3); 
    private readonly Semaphore second = new Semaphore(0, 2);
    public Foo() {
    }

    public void First(Action printFirst) {
        // printFirst() outputs "first". Do not change or remove this line.
        printFirst();
        first.Release(3);
    }

    public void Second(Action printSecond) {
        first.WaitOne();
        // printSecond() outputs "second". Do not change or remove this line.
        printSecond();
        second.Release(2);
    }

    public void Third(Action printThird) {
        first.WaitOne();
        second.WaitOne();
        // printThird() outputs "third". Do not change or remove this line.
        printThird();
    }
}