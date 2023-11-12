import axios, { AxiosError } from "axios";
import { Pagination } from "../classes/pagination";
import { Role } from "./roles-api";


const endpoint = "user"
const constToken = "userToken";

export async function GetUsers(pageNumber: number, pageSize: number, searchParam: string): Promise<Pagination<User> | AxiosError>{

    let url = `${process.env.REACT_APP_BASE_URL}/api/${endpoint}`
    const headers = {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${sessionStorage.getItem(constToken)}`,
      };

    url = `${url}?pageNumber=${pageNumber}&pageSize=${pageSize}`;

    if(searchParam) {
        url = `${url}&search=${searchParam}`;
    }
    
    try{
        var result = await axios.get<Pagination<User>>(url,{headers});
        return result.data;
    }
    catch (error){
        return error as AxiosError
    }
}

export async function GetUser(userId:string): Promise<User | AxiosError>{
    let url = `${process.env.REACT_APP_BASE_URL}/api/${endpoint}/${userId}`
    const headers = {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${sessionStorage.getItem(constToken)}`,
      };
    
    try{
        var result = await axios.get<User>(url,{headers});
        return result.data;
    }
    catch (error){
        return error as AxiosError
    }
}

export class User {
    Id!:        string;
    Email!:     string;
    Username!:  string;
    Firstname!: string;
    Lastname!:  string;
    Roles!:     Role[];
}





