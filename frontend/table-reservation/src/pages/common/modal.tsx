interface IModal {
  handleClose: () => void;
  show: boolean;
  children: React.ReactElement;
  title: string;
}

export function Modal({
  handleClose,
  show,
  children,
  title,
}: IModal): React.ReactElement {
  return (
    <div className={`modal ${show ? "display-block" : "display-none"}`}>
      <div className="modal-back-content" onClick={handleClose}></div>
      <section className="modal-main">
        <div className="modal-header">
          <h1>{title}</h1>
        </div>
        <div className="modal-content">{children}</div>
      </section>
    </div>
  );
}
