# Szurubooru React Frontend

This is a modern, standalone frontend for the Szurubooru image board engine, built with React and Vite.

## Features

*   User authentication (login/logout)
*   Browse and view posts with pagination
*   Search for posts by tags
*   View post details
*   Upload new posts

## Prerequisites

*   [Node.js](https://nodejs.org/) (version 16 or later recommended)
*   [npm](https://www.npmjs.com/)

## Getting Started

### 1. Installation

Navigate to the frontend directory and install the dependencies:

```bash
cd szurubooru-frontend
npm install
```

### 2. Running the Development Server

To start the development server, run the following command from the `szurubooru-frontend` directory:

```bash
npm run dev
```

This will start the application on `http://localhost:5173` by default.

**Note:** This frontend expects the Szurubooru backend API to be running and accessible. You may need to configure a proxy in `vite.config.js` to forward API requests to your backend server to avoid CORS issues.

Example `vite.config.js` with proxy:

```javascript
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080', // Your Szurubooru backend URL
        changeOrigin: true,
      },
    },
  },
})
```

### 3. Building for Production

To create a production build of the application, run:

```bash
npm run build
```

This will create a `dist` folder with the optimized, static assets for deployment.
