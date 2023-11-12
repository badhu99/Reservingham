import internal from "stream";
import { User } from "../../../apis/user-api";

interface ITableWrapper {
  pageNumber: number;
  pageSize: number;
  allItemsCount: number;
  nextPage: () => void;
  previousPage: () => void;
  search: (phrase: string) => void;
  children: React.ReactElement;
}

export function TableWrapper({
  pageNumber,
  pageSize,
  allItemsCount,
  nextPage,
  previousPage,
  search,
  children,
}: ITableWrapper) {
  const boolClickPreviousPage = pageNumber > 1;
  const boolClickNextPage = pageNumber * pageSize < allItemsCount;
  return (
    <div className="page-content">
      <section className="main-header grid">
        <input type="text" placeholder="Search..." />
        {/* <a className="button link">
            <span>Filters</span>
            <i className="fa-solid fa-angle-down"></i>
          </a> */}
        <button className="button">
          <i className="fa-solid fa-plus"></i>
          <span>Add new user</span>
        </button>
      </section>
      {children}
      <section className="table-footer grid">
        <span>
          Displaying {(pageNumber - 1) * pageSize + 1}-
          {boolClickNextPage ? pageNumber * pageSize : allItemsCount} of{" "}
          {allItemsCount} items
        </span>
        <div className="paging grid">
          <span>
            Page
            <input type="number" value={pageNumber} />
            of {Math.floor(allItemsCount / pageSize) + 1}
          </span>
          <div
            className="button icon"
            onClick={boolClickPreviousPage ? previousPage : undefined}
          >
            <i className="fa-solid fa-angle-left" />
          </div>
          <div
            className="button icon"
            onClick={boolClickNextPage ? nextPage : undefined}
          >
            <i className="fa-solid fa-angle-right" />
          </div>
        </div>
      </section>
    </div>
  );
}
