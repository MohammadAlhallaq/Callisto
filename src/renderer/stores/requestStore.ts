import { create } from 'zustand';
import type { HttpMethod, KeyValuePair, ResponseData, RequestError } from '../../shared/types';

interface RequestState {
  method: HttpMethod;
  url: string;
  headers: KeyValuePair[];
  params: KeyValuePair[];
  body: string;
  response: ResponseData | null;
  error: string | null;
  loading: boolean;

  setMethod: (method: HttpMethod) => void;
  setUrl: (url: string) => void;
  setHeaders: (headers: KeyValuePair[]) => void;
  setParams: (params: KeyValuePair[]) => void;
  setBody: (body: string) => void;
  sendRequest: () => Promise<void>;
}

export const useRequestStore = create<RequestState>((set, get) => ({
  method: 'GET',
  url: '',
  headers: [{ key: '', value: '', enabled: true }],
  params: [{ key: '', value: '', enabled: true }],
  body: '',
  response: null,
  error: null,
  loading: false,

  setMethod: (method) => set({ method }),
  setUrl: (url) => set({ url }),
  setHeaders: (headers) => set({ headers }),
  setParams: (params) => set({ params }),
  setBody: (body) => set({ body }),

  sendRequest: async () => {
    const { method, url, headers, params, body } = get();
    if (!url) return;

    set({ loading: true, error: null, response: null });

    try {
      const result = await window.callistoApi.sendRequest({
        method,
        url,
        headers,
        params,
        body,
      });

      if ('status' in result) {
        set({ response: result, loading: false });
      } else {
        set({ error: (result as RequestError).message, loading: false });
      }
    } catch (err) {
      set({ error: 'Failed to send request', loading: false });
    }
  },
}));
