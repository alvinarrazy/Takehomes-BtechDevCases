import React, { useEffect, useState } from 'react';
import { login } from '../api';
import { useNavigate } from 'react-router';
import { useGetUser } from '../services';

interface LoginForm {
  email: string;
  password: string;
}

export default function Login() {
  const [form, setForm] = useState<LoginForm>({
    email: '',
    password: '',
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const { data: user } = useGetUser();

  const navigate = useNavigate();

  useEffect(() => {
    if (user) {
      console.log('User already logged in, redirecting to home.');
      navigate('/');
    }
  }, [user]);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]: value,
    }));
    if (error) setError(null);
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    try {
      await login(form.email, form.password);

      setForm({ email: '', password: '' });
    } catch (err: any) {
      if (err.response?.data?.message) {
        setError(err.response.data.message);
      }
      console.error('Login error:', err);
    } finally {
      setLoading(false);
      navigate('/');
    }
  };

  return (
    <div
      style={{
        maxWidth: '400px',
        margin: '50px auto',
        padding: '20px',
        fontFamily: 'Arial, sans-serif',
      }}>
      <h2 style={{ textAlign: 'center', marginBottom: '30px' }}>Login</h2>

      <form onSubmit={handleSubmit}>
        <div style={{ marginBottom: '15px' }}>
          <label
            htmlFor='email'
            style={{ display: 'block', marginBottom: '5px' }}>
            Email:
          </label>
          <input
            type='text'
            id='email'
            name='email'
            value={form.email}
            onChange={handleInputChange}
            required
            style={{
              width: '100%',
              padding: '8px',
              border: '1px solid #ddd',
              borderRadius: '4px',
              fontSize: '14px',
            }}
            placeholder='Enter username (try: admin@email.com)'
          />
        </div>

        <div style={{ marginBottom: '20px' }}>
          <label
            htmlFor='password'
            style={{ display: 'block', marginBottom: '5px' }}>
            Password:
          </label>
          <input
            type='password'
            id='password'
            name='password'
            value={form.password}
            onChange={handleInputChange}
            required
            style={{
              width: '100%',
              padding: '8px',
              border: '1px solid #ddd',
              borderRadius: '4px',
              fontSize: '14px',
            }}
            placeholder='Enter password (try: password123)'
          />
        </div>

        <button
          type='submit'
          disabled={loading}
          style={{
            width: '100%',
            padding: '10px',
            backgroundColor: loading ? '#ccc' : '#007bff',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            fontSize: '16px',
            cursor: loading ? 'not-allowed' : 'pointer',
          }}>
          {loading ? 'Logging in...' : 'Login'}
        </button>
      </form>

      {error && (
        <div
          style={{
            marginTop: '15px',
            padding: '10px',
            backgroundColor: '#f8d7da',
            color: '#721c24',
            border: '1px solid #f5c6cb',
            borderRadius: '4px',
          }}>
          {error}
        </div>
      )}
    </div>
  );
}
