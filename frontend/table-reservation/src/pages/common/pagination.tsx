import "./pagination.scss";

import React, { ReactNode } from "react";

const PaginationComponent = ({ children }: { children: React.ReactNode }) => {
  return (
    <>
      <p>Upper pagination</p>
      {children}
      <Pagination1></Pagination1>
      {/* <Pagination2></Pagination2> */}
    </>
  );
};

const Pagination1 = () => {
  return (
    <div className="container1">
      <ul className="pagination1">
        <li>
          <a href="#">Prev</a>
        </li>
        <li>
          <a href="#">1</a>
        </li>
        <li className="active1">
          <a href="#">2</a>
        </li>
        <li>
          <a href="#">3</a>
        </li>
        <li>
          <a href="#">4</a>
        </li>
        <li>
          <a href="#">5</a>
        </li>
        <li>
          <a href="#">Next</a>
        </li>
      </ul>
    </div>
  );
};

const Pagination2 = () => {
  return (
    <div id="app" className="container">
      <ul className="page">
        <li className="page__btn">
          <span className="material-icons">chevron_left</span>
        </li>
        <li className="page__numbers"> 1</li>
        <li className="page__numbers active">2</li>
        <li className="page__numbers">3</li>
        <li className="page__numbers">4</li>
        <li className="page__numbers">5</li>
        <li className="page__numbers">6</li>
        <li className="page__dots">...</li>
        <li className="page__numbers"> 10</li>
        <li className="page__btn">
          <span className="material-icons">chevron_right</span>
        </li>
      </ul>
    </div>
  );
};

export default PaginationComponent;
