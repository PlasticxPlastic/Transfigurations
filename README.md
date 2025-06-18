# Transfiguration

A web application for formula calculations built with Vue.js and Go.

## Project Structure
- `frontend/`: Vue.js frontend application
- `backend/`: Go backend application

## Prerequisites
- Go 1.21 or later
- Node.js 18 or later
- PostgreSQL 15 or later

## Setup Instructions

### Backend Setup
1. Navigate to the backend directory:
```bash
cd backend
```

2. Install Go dependencies:
```bash
go mod download
```

3. Set up PostgreSQL database:
- Create a database named 'transfiguration'
- Update database connection settings in backend/internal/database/config.go

4. Run the backend server:
```bash
go run cmd/main.go
```

### Frontend Setup
1. Install dependencies:
```bash
cd frontend
npm install
```

2. Run development server:
```bash
npm run dev
```

## Development
- Backend runs on http://localhost:8080
- Frontend runs on http://localhost:5173

## Deployment
This project is configured for deployment on Vercel. 