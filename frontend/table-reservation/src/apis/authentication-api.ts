import axios, { AxiosError } from "axios";
import { LoginResponseModel, LoginRequestModel } from "../classes/login_model";


const endpoint = "auth"

export async function loginUser(data: LoginRequestModel): Promise<LoginResponseModel | AxiosError>{

    const action = "signin"
    const url = `${process.env.REACT_APP_BASE_URL}/api/${endpoint}/${action}`

    const headers = {
        "Content-Type": "application/json",
      };
    
    try{
        var result = await axios.post<LoginResponseModel>(url, data, {headers});
        return result.data;
    }
    catch (error){
        return error as AxiosError
    }
}



