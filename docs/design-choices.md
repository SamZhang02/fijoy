# Design Choices

This is a document that I use to keep track of all the design choices I made,
and I hope this also helps contributors to gain insight on how and why things
are done in the way they are, and please feel free to challenge those design
decisions as certainly they are not perfect :)

## What I want to achieve with Fijoy

There are a couple things I would like to achive with Fijoy

- The account balance is traceable at all time. Meaning if an account has X amount
  of money, it is the sum of all the transactions.
  Everything always adds up at all time.
- A great manual transaction logging experience, I have no interest at the
  moment for providers like Plaid or Flink, as they are quite unreliable
  based on my own testing, and most importantly, I think they defeat the
  purpose of doing expense tracking, which is making you more aware of your
  spending.

---

## Accounts

- Each account will maintain its current balance and the currency.
- Every time when there is a transaction, the balance will be updated accordingly.
- To plot the account balance history, we will do a range query over
  the transaction time of all the transactions for this account.
- The account balance at all time is the sum of all the transactions.
  Therefore, everything is traceable.

## Debt vs. Asset

If I have $100 (asset), that will be displayed as $100, and will be stored as
$100 in the database. IF I own $100 (debt), it will STILL be displayed as $100
as the "credit balance" on the client side, but stored as -$100 in the database.

## Transactions

- Make sure to update the balance of the accounts involved in the transaction.

### Transfer between accounts with the same currency

Easy.

### Transfer between accounts with different currencies

Take the example from RMB to CAD.

If I transfer 1000 RMB to my CAD account, that will results in roughly 200 CAD.

We will store 2 separate transactions with the type `transfer`.

By default, we assume the transfer does not change one's net worth,
meaning there is neither loss nor gain after making the transfer.

To take account into the potential loss/gain, we will prompt the user to input
the total loss/gain. Take the same example, if I transfer 1000 RMB to my CAD
account and I paid a 2 CAD transfer fee, receiving 198 CAD.

This will result in 3 transactions:

1. -1000 RMB from the RMB account, type `transfer`
2. +200 CAD to the CAD account, type `transfer`
3. -2 CAD from the CAD account, type `expense`. In case of gain, it will be `income`

This way, the monthly expense/income calculation becomes straightforward.

## Transaction Group

I do not believe they are useful. Change my mind.

## Montly "Statement"

- Every transaction will also update the corresponding monthly statement.
- Which exchange rate should be used for each monthly statement? The latest one
  or the one from the last day of that month? "The last day of that month" makes
  more sense I think.

## Exchange Rate Caching

TBD.

## Investment

How to track stacks? Take a brokerage account for example. We will first let the
user create that account like any other account (input the total value, choose
the currency...). Then, we will also let the user describe the underlying holdings.

The holdings will be stored, giving us the ability to fetch the latest price of
the given brokerage account. However! We will not automatically update the value
of the investment account. To update the value, user can either manually input
the delta as an income/expense transaction. Or they can click a "checkpoint"
button where we can automatically calculate the delta and create a transaction
from the last checkpoint to the current value.

## Payer / Payee

If this is an income transaction, the payer will be stored, and for an
expense transaction, the payee will be stored. Don't store anything for transfer
since it is between the same person's accounts.
