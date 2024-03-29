import axios, { AxiosError } from "axios";
import { Pagination } from "../classes/pagination";
import { Company } from "../classes/company";


const endpoint = "company"
const constToken = "userToken";

export async function GetCompanies(pageNumber: number, pageSize: number): Promise<Pagination<Company> | AxiosError>{

    let url = `${process.env.REACT_APP_BASE_URL}/api/${endpoint}`
    const headers = {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${sessionStorage.getItem(constToken)}`,
      };

    url = `${url}?pageNumber=${pageNumber}&pageSize=${pageSize}`
    
    try{
        var result = await axios.get<Pagination<Company>>(url,{headers});
        return result.data;
    }
    catch (error){
        return error as AxiosError
    }
}



