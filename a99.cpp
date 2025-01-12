#include <iostream>
#include <stack>

int main() {
    std::stack<int> myStack;

    myStack.push(10);
    myStack.push(20);
    myStack.push(30);

    for (int i = 0; i < 5; ++i) {
        std::cout << "Top element: " << myStack.top() << std::endl;
        myStack.pop();
    }

    if (myStack.empty()) {
        std::cout << "The stack is empty. Cannot access top." << std::endl;
        std::cout << "Top element: " << myStack.top() << std::endl;
    }

    myStack.push("string");

    return 0;
}
