import {FC} from "react";
import {IDelivery} from "../models/Order";
import {Col, Row} from "antd";

type DeliveryDetailsProps = {
    details: IDelivery
}

export const DeliveryDetails: FC<DeliveryDetailsProps> = (props) => {

    return (
        <Row gutter={[8, 16]}>
            <Col span={12}>
                Name
            </Col>
            <Col span={12}>
                {props.details.name}
            </Col>

            <Col span={12}>
                Phone
            </Col>
            <Col span={12}>
                {props.details.phone}
            </Col>

            <Col span={12}>
                Zip
            </Col>
            <Col span={12}>
                {props.details.zip}
            </Col>

            <Col span={12}>
                City
            </Col>
            <Col span={12}>
                {props.details.city}
            </Col>

            <Col span={12}>
                Address
            </Col>
            <Col span={12}>
                {props.details.address}
            </Col>

            <Col span={12}>
                Region
            </Col>
            <Col span={12}>
                {props.details.region}
            </Col>

            <Col span={12}>
                Email
            </Col>
            <Col span={12}>
                {props.details.email}
            </Col>
        </Row>
    );
}