export class LoginRequestModel{
    Username!: string;
    Password!:string;
}

export class LoginResponseModel{
    Id!: string;
    Email!:string;
    RefreshToken!: string;
    Username!:string;
    AccessToken!:string;
}