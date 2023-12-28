import axios, { AxiosError } from "axios";
import { Pagination } from "../classes/pagination";

const endpoint = "draft";
const constToken = "userToken";
export async function GetDrafts(
  pageNumber: number,
  pageSize: number,
  searchParam: string
): Promise<Pagination<Draft> | AxiosError> {
  let url = `${process.env.REACT_APP_BASE_URL}/api/${endpoint}`;
  const headers = {
    "Content-Type": "application/json",
    Authorization: `Bearer ${sessionStorage.getItem(constToken)}`,
  };

  url = `${url}?pageNumber=${pageNumber}&pageSize=${pageSize}`;

  try {
    var result = await axios.get<Pagination<Draft>>(url, { headers });
    return result.data;
  } catch (error) {
    return error as AxiosError;
  }
}

export class Draft {
  Id!: string;
  Name!: string;
}
