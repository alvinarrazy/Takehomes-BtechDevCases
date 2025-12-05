import { useGetUser } from '../services';

export default function Profile() {
  const { data: userRes } = useGetUser();

  return <div>Email: {userRes?.data.email}</div>;
}
