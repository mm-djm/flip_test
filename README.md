# flip_test

Build the checkout
backend service that will support different promotions with the given inventory.

| Name           | Price     | Inventory |
| -------------- | --------- | --------- |
| Google Home    | $49.99    | 10        |
| MacBook Pro    | $5,399.99 | 5         |
| Alexa Speaker  | $109.50   | 10        |
| Raspberry Pi B | $30.00    | 2         |

### The system should have the following promotions:

> - Each sale of a MacBook Pro comes with a free Raspberry Pi B
> - Buy 3 Google Homes for the price of 2
> - Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers

## Test

Use the following command:

```
cd ./shell
sh ./test.sh
```
