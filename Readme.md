# Project Overview

The objective of this project is to create an online marketplace specifically for sustainable and eco-friendly products. The marketplace will connect eco-conscious consumers with vendors who offer environmentally friendly products, ensuring transparency and authenticity through rigorous verification processes. The platform will not only facilitate the buying and selling of products but also provide educational resources about sustainability.

# Use Cases

## Guests

1. As a Guest, I'd like to access to the list of products so I can view what's available in the marketplace.
1. As a Guest, I'd like to see Product details so I can confirm I'm buying the right one.
1. As a Guest, I'd like to register so I can perform products purchases. It's required to have a Name, Email, Username and a Password
1. As a Guest, I'd like to authenticate in the system when I have an account using my username/email and password.

## Customers

1. As a Customer, I'd like to purchase a product so I can satisfy my needs.

## Sellers

1. As a Seller, I'd like to create a products in the marketplace so Customers can buy them.
1. As a Seller, I'd like to manage my marketplace products so I can adjust them based on Customer need.
1. As a Seller, I'd like report of most seller products so I can take decision over those ones.

### Appendix

All use cases not defined:

    User Authentication and Profiles:
        Users (both buyers and sellers) should be able to register, log in, and manage their profiles.
        Sellers can create and manage their storefronts.

    Product Listings:
        Sellers can list products with detailed descriptions, images, pricing, and sustainability certifications.
        Each product listing should include information about the product's environmental impact.

    Search and Filter:
        Buyers can search for products using various filters such as category, price range, and sustainability certifications.
        Implement advanced search functionalities like keyword search and sorting by popularity, newest, and price.

    Shopping Cart and Checkout:
        Buyers can add products to a shopping cart and proceed to checkout.
        Support for multiple payment methods including credit cards, PayPal, and other payment gateways.

    Order Management:
        Buyers can view their order history and track current orders.
        Sellers can manage orders, update order statuses, and handle returns and exchanges.

    Reviews and Ratings:
        Buyers can leave reviews and ratings for products and sellers.
        Implement moderation to ensure reviews are constructive and appropriate.

    Sustainability Education:
        A section dedicated to articles, blogs, and resources about sustainable living and eco-friendly practices.
        Integration with social media for content sharing and community engagement.

    Analytics Dashboard:
        Sellers have access to a dashboard showing sales analytics, customer feedback, and inventory management tools.
        Admin dashboard for managing the platform, including user management, product verification, and platform statistics.

    Mobile Responsiveness:
        Ensure the platform is fully responsive and works well on both desktop and mobile devices.

    Security and Privacy:
        Implement robust security measures to protect user data.
        Ensure compliance with data protection regulations like GDPR.


# Technical Considerations

## How to generate mocks?

1. Run `mockgen <Full Package Name> <Sources> >> <Sources Location Path>/mocks/mocks.go`

    Example:
    mockgen github.com/ezekielriva/ecommerce_go/src/core/repositories ProductRepository,UserRepository > src/core/repositories/mocks/mocks.go
