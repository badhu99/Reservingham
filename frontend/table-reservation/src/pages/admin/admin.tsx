import React, { useEffect, useState } from "react";
import AdminItem from "./components/single-item-admin";
import { GetCompanies } from "../../apis/company-api";
import { AxiosError } from "axios";
import { Pagination } from "../../classes/pagination";
import { Company } from "../../classes/company";
import PaginationWrapper from "./components/pagination-wrapper";
import { Modal } from "../common/modal";
import { TableWrapper } from "../manager/components/table-wrapper";

export default function Admin() {
  const [companyData, setCompanyData] = useState<Pagination<Company>>();
  const [pageNumber, setPageNumber] = useState(1);
  const [pageSize, setPageSize] = useState(5);
  const [search, setSearch] = useState("")
  const [modalCreateNew, setModalCreateNew] = useState(false);

  useEffect(() => {
    getData();
  }, [pageNumber, pageSize]);

  const getData = async () => {
    var data = await GetCompanies(pageNumber, pageSize);
    if (!(data instanceof AxiosError)) {
      setCompanyData(data);
    }
  };

  const nextPage = () => {
    setPageNumber((p) => p + 1);
  };

  const previousPage = () => {
    setPageNumber((p) => p - 1);
  };

  const toggleModalCreateNew = () => {
    setModalCreateNew((prev) => !prev);
  };

  const searchBy = (phrase: string) => {
    console.log("Admin TODO search")
  }

  return (
    <div className="admin-page">
      <Modal show={modalCreateNew} handleClose={toggleModalCreateNew} title="Admin">
        <p>Hello world</p>
      </Modal>
      <header className="header">
        <h1>Admin</h1>
      </header>
      {/* <div className="admin-content">
        <PaginationWrapper
          pageNumber={pageNumber}
          pageSize={companyData?.PageSize!}
          countItems={companyData?.Count!}
          onNextClick={nextPage}
          onPreviousClick={previousPage}
          openModal={toggleModalCreateNew}
        >
          <table className="table-data">
            <thead>
              <tr>
                <th>#</th>
                <th>Name</th>
                <th>Person num</th>
                <th></th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              {companyData &&
                companyData.Items.map((item, index) => {
                  return (
                    <AdminItem
                      id={index + 1 + (pageNumber - 1) * pageSize}
                      name={item.Name}
                      num={15}
                    />
                  );
                })}
            </tbody>
          </table>
        </PaginationWrapper>
      </div> */}
            {companyData && (
        <TableWrapper
          pageNumber={pageNumber}
          pageSize={pageSize}
          allItemsCount={companyData.Count}
          previousPage={previousPage}
          nextPage={nextPage}
          search={searchBy}
          openModal={() => {}}
          children={<CompanyTableData data={companyData.Items}/>}
        />
      )}
    </div>
  );
}

const CompanyTableData = ({ data }: { data: Company[] }) => {
  return (
    <div className="table-content">
      <table>
      <thead>
          <tr>
            <th></th>
            <th>ID</th>
            <th>Name</th>
          </tr>
        </thead>

        <tbody>
          {data &&
            data.map((item) => {
              return (
                <tr className="selected">
                  <td></td>
                  <td>1</td>
                  <td>{item.Name}</td>
                </tr>
              );
            })}
        </tbody>
      </table>
    </div>
  );
};