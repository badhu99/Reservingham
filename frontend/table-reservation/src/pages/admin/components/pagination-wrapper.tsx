import Pagination from "../../common/pagination";

const PaginationWrapper = ({
  children,
  pageNumber,
  pageSize,
  countItems,
  onNextClick,
  onPreviousClick,
  openModal
}: {
  children: React.ReactNode;
  pageNumber: number;
  pageSize: number;
  countItems: number;
  onNextClick: () => void;
  onPreviousClick: () => void;
  openModal: () => void;
}) => {
  return (
    <>
      <div className="upper">
        <div className="upper-left">
          <input type="text" placeholder="Search..."/>
        </div>
        <div className="upper-right">
          <button onClick={openModal}>Create</button>
        </div>
      </div>
      {children}
      <div className="bottom">
        <div className="bottom-left">
          <Pagination
            pageNumber={pageNumber}
            pageSize={pageSize}
            countItems={countItems}
            onNextClick={onNextClick}
            onPreviousClick={onPreviousClick}
          ></Pagination>
        </div>
        <div className="bottom-right">
          <p>Showing {(pageNumber - 1)  * pageSize + 1} - {(pageNumber * pageSize > countItems) ? countItems :pageNumber * pageSize} out of {countItems}</p>
        </div>
      </div>
    </>
  );
};

export default PaginationWrapper;
