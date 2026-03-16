import { contextBridge, ipcRenderer } from 'electron';
import type { RequestConfig, ResponseData, RequestError } from './shared/types';

const api = {
  sendRequest: (config: RequestConfig): Promise<ResponseData | RequestError> => {
    return ipcRenderer.invoke('send-request', config);
  },
};

contextBridge.exposeInMainWorld('callistoApi', api);

export type CallistoApi = typeof api;
