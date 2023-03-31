import {FC} from "react";
import {IPayment} from "../models/Order";
import {Col, Row} from "antd";

type PaymentDetailsProps = {
    details: IPayment
}

export const PaymentDetails: FC<PaymentDetailsProps> = (props) => {

    return(
        <Row gutter={[8, 16]}>
            <Col span={12}>
                Transaction
            </Col>
            <Col span={12}>
                {props.details.transaction}
            </Col>

            <Col span={12}>
                Request id
            </Col>
            <Col span={12}>
                {props.details.request_id}
            </Col>

            <Col span={12}>
                Currency
            </Col>
            <Col span={12}>
                {props.details.currency}
            </Col>

            <Col span={12}>
                Provider
            </Col>
            <Col span={12}>
                {props.details.provider}
            </Col>

            <Col span={12}>
                Amount
            </Col>
            <Col span={12}>
                {props.details.amount}
            </Col>

            <Col span={12}>
                Payment dt
            </Col>
            <Col span={12}>
                {props.details.payment_dt}
            </Col>

            <Col span={12}>
                Bank
            </Col>
            <Col span={12}>
                {props.details.bank}
            </Col>

            <Col span={12}>
                Delivery cost
            </Col>
            <Col span={12}>
                {props.details.delivery_cost}
            </Col>

            <Col span={12}>
                Goods total
            </Col>
            <Col span={12}>
                {props.details.goods_total}
            </Col>

            <Col span={12}>
                Custom fee
            </Col>
            <Col span={12}>
                {props.details.custom_fee}
            </Col>
        </Row>
    );
}
