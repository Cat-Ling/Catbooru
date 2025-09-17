import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { createPost } from '../services/api';

const UploadPage = () => {
  const [file, setFile] = useState(null);
  const [tags, setTags] = useState('');
  const [safety, setSafety] = useState('safe');
  const [source, setSource] = useState('');
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!file) {
      setError('Please select a file to upload.');
      return;
    }
    setError(null);
    setIsLoading(true);

    try {
      const newPost = await createPost({ file, tags, safety, source });
      setIsLoading(false);
      // Redirect to the new post's detail page
      navigate(`/post/${newPost.id}`);
    } catch (err) {
      setError('Upload failed. Please try again.');
      setIsLoading(false);
    }
  };

  return (
    <div>
      <h2>Upload New Post</h2>
      <form onSubmit={handleSubmit}>
        <div style={{ marginBottom: '1rem' }}>
          <label htmlFor="file">File:</label><br />
          <input id="file" type="file" onChange={handleFileChange} required />
        </div>
        <div style={{ marginBottom: '1rem' }}>
          <label htmlFor="tags">Tags (space-separated):</label><br />
          <input id="tags" type="text" value={tags} onChange={(e) => setTags(e.target.value)} style={{ width: '300px' }} />
        </div>
        <div style={{ marginBottom: '1rem' }}>
          <label htmlFor="safety">Safety:</label><br />
          <select id="safety" value={safety} onChange={(e) => setSafety(e.target.value)}>
            <option value="safe">Safe</option>
            <option value="sketchy">Sketchy</option>
            <option value="unsafe">Unsafe</option>
          </select>
        </div>
        <div style={{ marginBottom: '1rem' }}>
          <label htmlFor="source">Source URL (optional):</label><br />
          <input id="source" type="url" value={source} onChange={(e) => setSource(e.target.value)} style={{ width: '300px' }} />
        </div>
        <button type="submit" disabled={isLoading}>
          {isLoading ? 'Uploading...' : 'Upload'}
        </button>
      </form>
      {error && <p style={{ color: 'red', marginTop: '1rem' }}>{error}</p>}
    </div>
  );
};

export default UploadPage;
