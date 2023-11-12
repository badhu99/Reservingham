import axios, { AxiosError } from "axios";
import { Pagination } from "../classes/pagination";


const endpoint = "role"
const constToken = "userToken";
export async function GetRoles(): Promise<Pagination<Role> | AxiosError>{

    let url = `${process.env.REACT_APP_BASE_URL}/api/${endpoint}`
    const headers = {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${sessionStorage.getItem(constToken)}`,
      };

    try{
        var result = await axios.get<Pagination<Role>>(url,{headers});
        return result.data;
    }
    catch (error){
        return error as AxiosError
    }
}

export class Role {
    Id!:   string;
    Name!: string;
}