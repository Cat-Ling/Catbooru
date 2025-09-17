import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link, useNavigate } from 'react-router-dom';
import LoginPage from './pages/LoginPage';
import PostListPage from './pages/PostListPage';
import PostDetailPage from './pages/PostDetailPage';
import UploadPage from './pages/UploadPage'; // Import the new component
import PrivateRoute from './components/PrivateRoute';
import './App.css';

const AppLayout = ({ children }) => {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem('authToken');
    localStorage.removeItem('username');
    navigate('/login');
  };

  return (
    <div>
      <nav style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '1rem', background: '#eee' }}>
        <div>
          <Link to="/" style={{ marginRight: '1rem', textDecoration: 'none' }}>Home</Link>
          <Link to="/upload" style={{ textDecoration: 'none' }}>Upload</Link>
        </div>
        <button onClick={handleLogout}>Logout</button>
      </nav>
      <main style={{ padding: '1rem' }}>
        {children}
      </main>
    </div>
  );
};

function App() {
  return (
    <Router>
      <div className="App">
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route
            path="/*"
            element={
              <PrivateRoute>
                <AppLayout>
                  <Routes>
                    <Route path="/" element={<PostListPage />} />
                    <Route path="/post/:id" element={<PostDetailPage />} />
                    <Route path="/upload" element={<UploadPage />} /> {/* Add the new route */}
                  </Routes>
                </AppLayout>
              </PrivateRoute>
            }
          />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
