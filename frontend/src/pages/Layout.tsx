import { Outlet, Link } from "react-router-dom";
import Nav from "react-bootstrap/Nav";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

const Layout = () => {
  return (
    <Container className="mt-2">
      <Row>
        <Nav variant="pills" activeKey="1">
          <Nav.Item>
            <Nav.Link eventKey="1" href="/" title="Home">
              Home
            </Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link eventKey="2" href="/users" title="Users">
              Users
            </Nav.Link>
          </Nav.Item>
        </Nav>
      </Row>
      <Row className="mt-2">
        <Outlet />
      </Row>
    </Container>
  );
};

export default Layout;
