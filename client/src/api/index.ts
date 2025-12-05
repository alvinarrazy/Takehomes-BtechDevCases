import axios, { type InternalAxiosRequestConfig } from 'axios';
import type { LoginResponse } from './types';

const api = (options?: Partial<InternalAxiosRequestConfig>) => {
  const axiosApi = axios.create();
  axiosApi.interceptors.request.use(
    async (config) => {
      config = {
        ...config,
        withCredentials: true,
        baseURL: 'http://localhost:8080',
        ...options,
      };

      return config;
    },
    (error) => {
      Promise.reject(error);
    },
  );

  axiosApi.interceptors.response.use(
    (response) => {
      response.data = response.data?.data || response.data;
      return response;
    },
    // async (error: AxiosError) => {
    async () => {
      let customError;

      //Handle error

      return Promise.reject(customError);
    },
  );
  return axiosApi;
};

export function login(email: string, password: string) {
  return api().post<LoginResponse>('/auth/login', {
    email: email,
    password: password,
  });
}

export function logout() {
  return api().post<LoginResponse>('/auth/logout');
}

export function getEmail() {
  return api().get<LoginResponse>('/user');
}
