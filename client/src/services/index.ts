import { useQuery } from '@tanstack/react-query';
import { getEmail } from '../api';

export function useGetUser() {
  return useQuery({ queryKey: ['user'], queryFn: getEmail });
}
