import React, { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import { getPost } from '../services/api';

const PostDetailPage = () => {
  const { id } = useParams();
  const [post, setPost] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchPost = async () => {
      setLoading(true);
      setError(null);
      try {
        const data = await getPost(id);
        setPost(data);
      } catch (err) {
        setError('Failed to load post details.');
      } finally {
        setLoading(false);
      }
    };

    fetchPost();
  }, [id]);

  if (loading) return <p>Loading post...</p>;
  if (error) return <p style={{ color: 'red' }}>{error}</p>;
  if (!post) return <p>Post not found.</p>;

  return (
    <div>
      <Link to="/">&larr; Back to all posts</Link>
      <h2>Post #{post.id}</h2>
      <div className="post-content" style={{ margin: '20px 0' }}>
        {post.type === 'image' && <img src={post.contentUrl} alt={`Post ${post.id}`} style={{ maxWidth: '100%', maxHeight: '80vh' }} />}
        {post.type === 'video' && <video src={post.contentUrl} controls style={{ maxWidth: '100%', maxHeight: '80vh' }} />}
        {/* TODO: Add handlers for other types like 'animation' if needed */}
      </div>
      <h3>Tags</h3>
      <div className="tags" style={{ display: 'flex', flexWrap: 'wrap', gap: '5px', marginBottom: '20px' }}>
        {post.tags.map(tag => (
          <span key={tag.names[0]} className="tag" style={{ background: '#ddd', padding: '2px 8px', borderRadius: '3px' }}>
            {tag.names[0]}
          </span>
        ))}
      </div>
      <h3>Details</h3>
      <ul>
        <li>Uploader: {post.user.name}</li>
        <li>Created at: {new Date(post.creationTime).toLocaleString()}</li>
        <li>Safety: {post.safety}</li>
        {post.source && <li>Source: <a href={post.source} target="_blank" rel="noopener noreferrer">{post.source}</a></li>}
        <li>Dimensions: {post.canvasWidth}x{post.canvasHeight}</li>
        <li>Score: {post.score}</li>
      </ul>
    </div>
  );
};

export default PostDetailPage;
