import { Link, Outlet, useNavigate } from 'react-router';
import { useGetUser } from '../services';
import { useEffect } from 'react';
import { logout } from '../api';
import { useQueryClient } from '@tanstack/react-query';

export default function ProtectedLayout() {
  const { isLoading, error } = useGetUser();
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  useEffect(() => {
    // Assume all error from getting user data is authentication related
    console.log('ProtectedLayout useEffect triggered with error:', error);
    if (error) {
      // Clear all cached data when authentication fails
      queryClient.resetQueries({ queryKey: ['user'] });
      navigate('/login');
    }
  }, [error, navigate, queryClient]);

  async function handleLogout() {
    try {
      await logout();
      // Clear all cached data on logout
      queryClient.resetQueries({ queryKey: ['user'] });
      navigate('/login');
    } catch (err) {
      console.error('Logout error:', err);
    }
  }

  return (
    <div>
      Secured Area
      <div>--------Other page--------</div>
      {isLoading ? (
        <div>Loading...</div>
      ) : error ? (
        <div>Error fetching user data</div>
      ) : (
        <>
          <Link to='/'>Home</Link> | <Link to='/profile'>Profile</Link>
          <Outlet />
        </>
      )}
      <button onClick={handleLogout}>Logout</button>
    </div>
  );
}
