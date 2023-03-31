import {FC} from "react";
import {Col, Row} from "antd";
import {IOrder} from "../models/Order";

type OrderDetailsProps = {
    details: IOrder
}

export const OrderDetails: FC<OrderDetailsProps> = (props) => {

    return (
        <Row gutter={[8, 16]}>
            <Col span={12}>
                Track number
            </Col>
            <Col span={12}>
                {props.details.track_number}
            </Col>

            <Col span={12}>
                Entry
            </Col>
            <Col span={12}>
                {props.details.entry}
            </Col>

            <Col span={12}>
                Locale
            </Col>
            <Col span={12}>
                {props.details.locale}
            </Col>

            <Col span={12}>
                Internal signature
            </Col>
            <Col span={12}>
                {props.details.internal_signature}
            </Col>

            <Col span={12}>
                Customer id
            </Col>
            <Col span={12}>
                {props.details.customer_id}
            </Col>

            <Col span={12}>
                Delivery service
            </Col>
            <Col span={12}>
                {props.details.delivery_service}
            </Col>

            <Col span={12}>
                Shard key
            </Col>
            <Col span={12}>
                {props.details.shardkey}
            </Col>

            <Col span={12}>
                sm id
            </Col>
            <Col span={12}>
                {props.details.sm_id}
            </Col>

            <Col span={12}>
                Date created
            </Col>
            <Col span={12}>
                {props.details.date_created}
            </Col>

            <Col span={12}>
                oof shard
            </Col>
            <Col span={12}>
                {props.details.oof_shard}
            </Col>
        </Row>
    );
}