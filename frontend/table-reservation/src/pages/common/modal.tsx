export function Modal({
  handleClose,
  show,
  children,
}: {
  handleClose: () => void;
  show: boolean;
  children: React.ReactElement;
}): React.ReactElement {

  return (
    <div className={`modal ${show ? "display-block" : "display-none"}`}>
      <div className="modal-back-content" onClick={handleClose}></div>
      <section className="modal-main">
        {children}
        <button type="button" onClick={handleClose}>
          Close
        </button>
      </section>
    </div>
  );
}
