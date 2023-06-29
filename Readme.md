## CashLink RestFull API

E-Wallet Restfull API using GO, GORM, PostgreSQL, and GIN.

---

## Description

This project is an e-wallet application that includes authentication features for two entities: users and merchants. Users and merchants can register an account, login, and update their profiles. Merchants can also create products that can be purchased by registered customers or users.

The transaction features in this application include top-up, payment, and balance transfer between users. Users need to log in to access these features. The application uses JWT authentication, ensuring the security of each user's account.

---

## Feature :

1. Accont Management
   1. User Account
   2. Merchant Account
2. Transfer
3. Payment
4. Topup
5. Transaction History

---

## Packages:

1. [GIN](https://gin-gonic.com/)
2. [GORM](https://gorm.io/)
3. [bycrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
4. [golang-jwt](https://github.com/golang-jwt/jwt)
5. [postgresql](https://www.postgresql.org/)

---

## API Endpoint

1. Account Management:

   - User :
     - Signup
       ```
       METHOD: POST
       /signup/user
       ```
     - Login
       ```
       METHOD: POST
       /login/user
       ```
     - Logout
       ```
       METHOD: POST
       /user/logout
       ```
     - View Profile
       ```
       METHOD: GET
       /user/:userID
       ```
     - Update Profile
       ```
       METHOD: PUT
       /user/:userID
       ```
     - Find User
       ```
       METHOD: GET
       /user/search/:userID/:username
       ```
   - Merchant :
     - Signup
       ```
       METHOD: POST
       /signup/merchant
       ```
     - Login
       ```
       METHOD: POST
       /login/merchant
       ```
     - Logout
       ```
       METHOD: POST
       /merchant/logout
       ```
     - View Merchant Profile
       ```
       METHOD: GET
       /merchant/:merchantID
       ```
     - Add New Product
       ```
       METHOD: POST
       /merchant/product/:merchantID
       ```
     - Update Product
       ```
       METHOD: PUT
       /merchant/product/:merchantID
       ```
     - View Product
       ```
       METHOD: GET
       /merchant/product/:merchantID
       ```
   - Transaction :

     **Require User Login**

     - Transfer
       ```
       METHOD: POST
       user/transaction/transfer/:userID/:recipientID
       ```
     - Payment
       ```
       METHOD: POST
       user/transaction/payment/:userID/:merchantID/:productID
       ```
     - Topup
       ```
       METHOD: POST
       user/transaction/topup/:userID
       ```
     - History Transaction
       ```
       METHOD: GET
       user/transaction/history/:userID
       ```
