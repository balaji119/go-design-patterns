**Problem without abstract factory**
- Adding a new payment method means modifying CreatePaymentProcessor().
- All creation logic is stuffed in one big function.
- Hard to test creation logic independently.
- Violates Open/Closed Principle.