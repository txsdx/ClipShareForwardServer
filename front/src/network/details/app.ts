import {request} from '../request'
import {AppVersion} from "@/types";

export const getVersion = () => {
  return request({
    url: '/version',
    method: 'get',
  }) as unknown as Promise<AppVersion>
}
