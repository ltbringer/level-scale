# 📝 Product Requirements Document

**Title**: E-commerce Platform – Buyer & Seller APIs (Level 1 MVP)  
**Version**: v1.0  
**Owner**: amresh.venugopal  
**Last Updated**: 2025-05-26

---

## 🎯 Purpose

This document defines the core functionality for the MVP version of a scalable e-commerce platform 
where users can act as **buyers**, **sellers**, or **both**. It outlines the expected user roles, 
authentication flows, and critical API-driven interactions that must be supported.

---

## 👤 User Roles

### 1. Buyer
- register and authenticate
- browse/search products in shops
- add items to a cart
- place orders and complete purchases
- Receives an invoice after ordering
- Receives a delivery schedule
- return purchased items
- rate and review products

### 2. Seller
- register and authenticate
- create and manage their shop
- list and manage products (with inventory, price, description)

### 3. Dual Role
- A single user can act as both buyer and seller

---

## 🔐 Authentication & Authorization

- Users must be able to register with email, password, and select role(s)
- Authentication via email/password and token (JWT or session-based)
- Role-based access controls for endpoints (e.g., only sellers can create products)

---

## 📦 Core Features & API Requirements

### 🧑 User APIs
- `POST /register` – Register user (buyer/seller/both)
- `POST /login` – Login and get auth token
- `GET /me` – Fetch user profile and roles

### 🏬 Shop & Product Management (Seller Only)
- `POST /shops` – Create seller shop
- `GET /shops` – List all shops
- `POST /products` – Add product to seller’s shop
- `PUT /products/:id` – Update product details
- `GET /products?shop_id=X` – List/search products by shop

### 🛒 Cart & Checkout (Buyer Only)
- `POST /cart` – Create or get current cart
- `POST /cart/items` – Add item to cart
- `GET /cart` – View cart contents
- `DELETE /cart/items/:id` – Remove item from cart
- `POST /checkout` – Place order based on cart

### 🧾 Post-Order (System Generated + Buyer Actions)
- `GET /orders/:id/invoice` – Get invoice for order
- `GET /orders/:id/delivery` – Get delivery schedule
- `POST /returns` – Request return for item
- `POST /products/:id/review` – Submit review + rating
- `GET /products/:id/reviews` – List reviews for product

---

## 📑 Additional Notes

- The system should auto-generate invoices upon order placement
- Delivery schedules can be simulated for now (e.g., +3 days)
- Reviews must be linked to verified purchases
- Sellers cannot modify reviews of their own products
