import axios from 'axios';

const API_URL = '/api';

/**
 * Logs in a user and returns a user token.
 * @param {string} username
 * @param {string} password
 * @returns {Promise<object>} The user token object from the API.
 */
export const login = async (username, password) => {
  try {
    const response = await axios.post(
      `${API_URL}/user-token/${username}`,
      {},
      {
        auth: {
          username,
          password,
        },
      }
    );
    // The API doc for "User token" resource specifies a "token" field.
    if (response.data && response.data.token) {
      return response.data;
    } else {
      throw new Error('Token not found in login response');
    }
  } catch (error) {
    console.error('Login failed:', error);
    throw error;
  }
};

const getAuthHeaders = () => {
  const token = localStorage.getItem('authToken');
  const username = localStorage.getItem('username');
  if (token && username) {
    // The API expects the token to be in the format 'username:token' and base64 encoded.
    const encoded = btoa(`${username}:${token}`);
    return {
      Authorization: `Token ${encoded}`,
    };
  }
  return {};
};

/**
 * Fetches a paginated list of posts, with an optional search query.
 * @param {number} offset - The starting point of the posts to fetch.
 * @param {number} limit - The number of posts to fetch.
 * @param {string} query - The search query string.
 * @returns {Promise<object>} The paginated list of posts from the API.
 */
export const getPosts = async (offset = 0, limit = 20, query = '') => {
  try {
    const params = { offset, limit };
    if (query && query.trim() !== '') {
      params.query = query;
    }
    const response = await axios.get(`${API_URL}/posts`, {
      params,
      headers: getAuthHeaders(),
    });
    return response.data;
  } catch (error) {
    console.error('Failed to fetch posts:', error);
    throw error;
  }
};

/**
 * Fetches a single post by its ID.
 * @param {number} id - The ID of the post to fetch.
 * @returns {Promise<object>} The post object from the API.
 */
export const getPost = async (id) => {
  try {
    const response = await axios.get(`${API_URL}/post/${id}`, {
      headers: getAuthHeaders(),
    });
    return response.data;
  } catch (error) {
    console.error(`Failed to fetch post with id ${id}:`, error);
    throw error;
  }
};

/**
 * Creates a new post by uploading a file and metadata.
 * @param {object} postData - The data for the new post.
 * @param {File} postData.file - The file to upload.
 * @param {string} postData.tags - A space-separated string of tags.
 * @param {string} postData.safety - The safety level ('safe', 'sketchy', 'unsafe').
 * @param {string} postData.source - The optional source URL.
 * @returns {Promise<object>} The newly created post object.
 */
export const createPost = async (postData) => {
  const { file, tags, safety, source } = postData;
  const formData = new FormData();

  const metadata = {
    tags: tags.split(' ').filter(t => t), // Split by space and remove empty strings
    safety,
    source,
  };

  formData.append('metadata', JSON.stringify(metadata));
  formData.append('content', file);

  try {
    const response = await axios.post(`${API_URL}/posts/`, formData, {
      headers: {
        ...getAuthHeaders(),
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  } catch (error) {
    console.error('Failed to create post:', error);
    throw error;
  }
};
