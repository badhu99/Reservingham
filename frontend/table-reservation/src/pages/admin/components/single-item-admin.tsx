const AdminItem = (props: any) => {
  return (
      <tr>
        <td>{props.id}</td>
        <td>{props.name}</td>
        <td>{props.num}</td>
        <td>Edit me</td>
        <td>Delete me</td>
      </tr>
  );
};

export default AdminItem;
