import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom'; // Import Link
import { getPosts } from '../services/api';

const PostListPage = () => {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [page, setPage] = useState(0);
  const [total, setTotal] = useState(0);
  const [searchQuery, setSearchQuery] = useState('');
  const [submittedQuery, setSubmittedQuery] = useState('');
  const limit = 20;

  useEffect(() => {
    const fetchPosts = async () => {
      setLoading(true);
      setError(null);
      try {
        const data = await getPosts(page * limit, limit, submittedQuery);
        setPosts(data.results);
        setTotal(data.total);
      } catch (err) {
        setError('Failed to load posts.');
      } finally {
        setLoading(false);
      }
    };

    fetchPosts();
  }, [page, submittedQuery]);

  const handleSearch = (e) => {
    e.preventDefault();
    setPage(0);
    setSubmittedQuery(searchQuery);
  };

  const handleNextPage = () => {
    if ((page + 1) * limit < total) {
      setPage(page + 1);
    }
  };

  const handlePrevPage = () => {
    if (page > 0) {
      setPage(page - 1);
    }
  };

  return (
    <div>
      <h2>Posts</h2>
      <form onSubmit={handleSearch} style={{ marginBottom: '20px' }}>
        <input
          type="text"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          placeholder="Search by tags..."
          style={{ marginRight: '10px' }}
        />
        <button type="submit">Search</button>
      </form>

      {loading && <p>Loading posts...</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}

      <div className="post-grid" style={{ display: 'flex', flexWrap: 'wrap', gap: '10px' }}>
        {posts.map((post) => (
          <div key={post.id} className="post-thumbnail">
            <Link to={`/post/${post.id}`}> {/* Wrap image in a Link */}
              <img src={post.thumbnailUrl} alt={`Post ${post.id}`} style={{ width: '150px', height: '150px', objectFit: 'cover', display: 'block' }} />
            </Link>
          </div>
        ))}
      </div>

      {!loading && posts.length === 0 && <p>No posts found.</p>}

      {total > 0 && (
        <div className="pagination" style={{ marginTop: '20px' }}>
          <button onClick={handlePrevPage} disabled={page === 0}>
            Previous
          </button>
          <span style={{ margin: '0 10px' }}>
            Page {page + 1} of {Math.ceil(total / limit)}
          </span>
          <button onClick={handleNextPage} disabled={(page + 1) * limit >= total}>
            Next
          </button>
        </div>
      )}
    </div>
  );
};

export default PostListPage;
