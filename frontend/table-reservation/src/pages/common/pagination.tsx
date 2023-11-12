const Pagination = ({
  pageNumber,
  pageSize,
  countItems,
  onNextClick,
  onPreviousClick,
}: {
  pageNumber: number;
  pageSize: number;
  countItems: number;
  onNextClick: () => void;
  onPreviousClick: () => void;
}) => {
  const checkPreviousClick = pageNumber - 1 > 0;
  const checkNextClick = countItems > pageNumber * pageSize;
  return (
    <>
      <ul className="pagination">
        {<li className={`page-buttons ${ checkPreviousClick ? "" : "hidden-button"}`} onClick={checkPreviousClick ? onPreviousClick : undefined}>
          Prev </li>}
        {<li className={`page-numbers ${checkPreviousClick ? "" : "hidden-button"}`} >
          {pageNumber - 1} </li> }
        <li className="page-numbers active">{pageNumber}</li>
        {<li className={`page-numbers ${checkNextClick ? "" : "hidden-button"}`}>
          {pageNumber + 1}</li> }
        {<li className={`page-numbers ${checkNextClick ? "" : "hidden-button"}`} onClick={checkNextClick ? onNextClick: undefined}> Next </li>}
      </ul>
    </>
  );
};

export default Pagination;
