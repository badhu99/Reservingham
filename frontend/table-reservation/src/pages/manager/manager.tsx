import { AxiosError } from "axios";
import { useState, useEffect } from "react";
import { User, GetUsers } from "../../apis/user-api";
import { Pagination } from "../../classes/pagination";
import { TableWrapper } from "./components/table-wrapper";
import { useNavigate } from "react-router-dom";
import { Modal } from "../common/modal";

export default function Manager() {
  const [pageNumber, setPageNumber] = useState(1);
  const [pageSize, setPageSize] = useState(10);
  const [searchParam, setSearchParam] = useState("");
  const [userData, setUserData] = useState<Pagination<User>>();
  const [openModal, setOpenModal] = useState(true);

  useEffect(() => {
    getUserData();
  }, [pageNumber, pageSize]);

  const getUserData = async () => {
    const data = await GetUsers(pageNumber, pageSize, searchParam);
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

  const toggleOpenModal = () => {
    setOpenModal((prev) => !prev);
  };

  return (
    <>
      <Modal handleClose={toggleOpenModal} show={openModal} title="Invite user">
        <CreateUser toggleModal={toggleOpenModal} />
      </Modal>
      <h1>User manager</h1>
      {userData && (
        <TableWrapper
          pageNumber={pageNumber}
          pageSize={pageSize}
          allItemsCount={userData.Count}
          previousPage={previousPage}
          nextPage={nextPage}
          search={search}
          openModal={toggleOpenModal}
          children={
            <UserTableData
              data={userData.Items}
              pageNumber={pageNumber}
              pageSize={pageSize}
            />
          }
        />
      )}
    </>
  );
}

const UserTableData = ({
  data,
  pageNumber,
  pageSize,
}: {
  data: User[];
  pageNumber: number;
  pageSize: number;
}) => {
  let navigate = useNavigate();

  const editUser = (user: User) => {
    navigate(`./${user.Id}`);
  };

  return (
    <div className="table-content">
      <table>
        <thead>
          <tr>
            <th></th>
            <th>ID</th>
            <th>Username</th>
            <th>First name</th>
            <th>Last name</th>
            <th>Email</th>
            <th></th>
          </tr>
        </thead>

        <tbody>
          {data &&
            data.map((item, index) => {
              return (
                <tr className="selected" key={item.Id}>
                  <td></td>
                  <td>{(pageNumber - 1) * pageSize + index + 1}</td>
                  <td>{item.Username}</td>
                  <td>{item.Firstname}</td>
                  <td>{item.Lastname}</td>
                  <td>{item.Email}</td>
                  <td>
                    <button
                      className="btn-git"
                      onClick={() => {
                        editUser(item);
                      }}
                    >
                      Edit
                    </button>
                  </td>
                </tr>
              );
            })}
        </tbody>
      </table>
    </div>
  );
};

interface ICreateUser {
  toggleModal: () => void;
}

const CreateUser: React.FC<ICreateUser> = ({ toggleModal }) => {
  return (
    <div className="create-user-content">
      <div className="div-main">
        <label>Email:</label>
        <input type="text" className="input-git" />
      </div>
      <div className="div-btn">
        <button className="btn-git">Invite</button>
        <button className="btn-git" onClick={toggleModal}>
          Close
        </button>
      </div>
    </div>
  );
};
