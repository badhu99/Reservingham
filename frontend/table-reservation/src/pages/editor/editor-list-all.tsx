import { AxiosError } from "axios";
import React, { useEffect, useState } from "react"
import { TableWrapper } from "../manager/components/table-wrapper";

import { Pagination } from "../../classes/pagination";
import { useNavigate } from "react-router-dom";
import { Draft, GetDrafts } from "../../apis/drafts-api";

export function EditorListAll(){

    const [pageNumber, setPageNumber] = useState(1);
    const [pageSize, setPageSize] = useState(10);
    const [searchParam, setSearchParam] = useState("");
    const [userData, setUserData] = useState<Pagination<Draft>>();
  
    useEffect(() => {
      getUserData();
    }, [pageNumber, pageSize]);
  
    const getUserData = async () => {
      const data = await GetDrafts(pageNumber, pageSize, searchParam);
      if (!(data instanceof AxiosError)) {
        setUserData(data);
      }
    };
  
    const nextPage = () => {
      setPageNumber((prev) => prev + 1);
    };
  
    const previousPage = () => {
      setPageNumber((prev) => prev - 1);
    };
  
    const search = (phrase: string) => {
      console.log(phrase);
    };

    return (
        <>
          <h1>Drafts</h1>
          {userData && (
            <TableWrapper
              pageNumber={pageNumber}
              pageSize={pageSize}
              allItemsCount={userData.Count}
              previousPage={previousPage}
              nextPage={nextPage}
              search={search}
              openModal={() => {}}
              children={<UserTableData data={userData.Items} pageNumber={pageNumber} pageSize={pageSize}/>}
            />
          )}
        </>
      );
}

const UserTableData = ({ data, pageNumber, pageSize }: { data: Draft[], pageNumber: number, pageSize: number }) => {

    let navigate = useNavigate()
  
    const editDraft = (draft:Draft) => {
      navigate(`./${draft.Id}`);
    }
      
    return (
      <div className="table-content">
        <table>
          <thead>
            <tr>
              <th></th>
              <th>ID</th>
              <th>Name</th>
              <th></th>
            </tr>
          </thead>
  
          <tbody>
            {data &&
              data.map((item, index) => {
                return (
                  <tr className="selected" key={item.Id}>
                    <td></td>
                    <td>{((pageNumber - 1) * pageSize) + index + 1}</td>
                    <td>{item.Name}</td>
                    <td>
                    <button className="btn-git" onClick={() => {editDraft(item)}}>Open</button>
                    <button className="btn-git" onClick={() => {editDraft(item)}}>Edit</button>
                        <button className="btn-git" onClick={() => {editDraft(item)}}>Delete</button>
                    </td>
                  </tr>
                );
              })}
          </tbody>
        </table>
      </div>
    );
  };