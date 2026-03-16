import type { CallistoApi } from '../preload';

declare global {
  interface Window {
    callistoApi: CallistoApi;
  }
}
