import { ipcMain } from 'electron';
import type { RequestConfig, ResponseData, RequestError } from '../../shared/types';

export function registerRequestHandlers() {
  ipcMain.handle('send-request', async (_event, config: RequestConfig): Promise<ResponseData | RequestError> => {
    const { method, url, headers, params, body } = config;

    // Build URL with query params
    const urlObj = new URL(url);
    for (const param of params) {
      if (param.enabled && param.key) {
        urlObj.searchParams.append(param.key, param.value);
      }
    }

    // Build headers
    const headersInit: Record<string, string> = {};
    for (const header of headers) {
      if (header.enabled && header.key) {
        headersInit[header.key] = header.value;
      }
    }

    const startTime = performance.now();

    try {
      const response = await fetch(urlObj.toString(), {
        method,
        headers: headersInit,
        body: method !== 'GET' && method !== 'HEAD' ? body || undefined : undefined,
      });

      const duration = Math.round(performance.now() - startTime);
      const responseBody = await response.text();

      const responseHeaders: Record<string, string> = {};
      response.headers.forEach((value, key) => {
        responseHeaders[key] = value;
      });

      return {
        status: response.status,
        statusText: response.statusText,
        headers: responseHeaders,
        body: responseBody,
        duration,
        size: new TextEncoder().encode(responseBody).length,
      };
    } catch (err) {
      return { message: err instanceof Error ? err.message : 'Request failed' };
    }
  });
}
