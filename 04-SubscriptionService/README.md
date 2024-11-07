# Subscription Service

A web-based subscription management service built with Go that handles user registration, authentication, and subscription plan management.

## Features

- User Authentication (Login/Logout)
- User Registration with Email Verification
- Subscription Plan Management
- Automated Invoice Generation
- PDF Manual Generation for Subscribers
- Email Notifications
- Session Management

## Technical Stack

- **Backend**: Go
- **Template Engine**: Go HTML Templates
- **PDF Generation**: gofpdf/gofpdi
- **Database**: PostgreSQL
- **Session Management**: Custom implementation

## Key Features Explained

### User Management

- User registration with email verification
- Secure login/logout functionality
- Account activation via email verification

### Subscription Management

- Multiple subscription plan support
- Plan selection and subscription processing
- Automated invoice generation
- Custom PDF manual generation for subscribers

### Email Notifications

- Registration confirmation emails
- Account activation emails
- Invoice delivery
- User manual delivery

## Setup and Installation

1. Clone the repository
2. Set up your PostgreSQL database
3. Configure your environment variables
4. Run the following commands:

```bash
go mod download
go run ./cmd/web
```

The service will be available at `http://localhost:8080`

## API Endpoints

- `GET /` - Home page
- `GET /login` - Login page
- `POST /login` - Process login
- `GET /logout` - Handle logout
- `GET /register` - Registration page
- `POST /register` - Process registration
- `GET /activate` - Account activation
- `GET /members/plans` - View subscription plans
- `GET /members/plan/subscribe` - Subscribe to a plan

## Security Features

- Session-based authentication
- Password hashing
- Email verification for new accounts
- Secure token generation for account activation
