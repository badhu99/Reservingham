import React, { useEffect, useState } from "react";
import AdminItem from "./components/single-item-admin";
import "./admin.scss";
import { GetCompanies } from "../../apis/company-api";
import { AxiosError } from "axios";
import { Pagination } from "../../classes/pagination";
import { Company } from "../../classes/company";
import { type } from "os";
import PaginationComponent from "../common/pagination";

const data = [
  { Id: 1, Name: "DeMarcus Cousins", Num: 10 },
  { Id: 2, Name: "DeMarcus Cousins1", Num: 11 },
  { Id: 3, Name: "DeMarcus Cousins12", Num: 4 },
  { Id: 3, Name: "DeMarcus Cousins134", Num: 5 },
  { Id: 3, Name: "DeMarcus Cousins433", Num: 6 },
  { Id: 3, Name: "DeMarcus Cousins1234", Num: 24 },
  { Id: 3, Name: "DeMarcus Cousins12344", Num: 222 },
];

export default function Admin() {
  const [companyData, setCompanyData] = useState<Pagination<Company>>();

  useEffect(() => {
    getData();
  }, [companyData]);

  const getData = async () => {
    var data = await GetCompanies();
    if (!(data instanceof AxiosError)) {
      setCompanyData(data);
    }
  };
  return (
    <div className="admin-page">
      <header className="header">
        <h1>Admin</h1>
      </header>
      <PaginationComponent>
        <table className="table-data">
          <thead>
            <tr>
              <th>#</th>
              <th>Name</th>
              <th>Person num</th>
              <th>Edit</th>
              <th>Delete</th>
            </tr>
          </thead>
          <tbody>
            {companyData &&
              companyData.Items.map((item, index) => {
                return <AdminItem id={index + 1} name={item.Name} num={15} />;
              })}
          </tbody>
        </table>
      </PaginationComponent>
    </div>
  );
}


