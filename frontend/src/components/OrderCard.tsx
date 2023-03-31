import React from "react";
import {IOrder} from "../models/Order";
import {Card, Divider} from "antd";
import {PaymentDetails} from "./PaymentDetails";
import {DeliveryDetails} from "./DeliveryDetails";
import {ItemDetails} from "./ItemDetails";
import {OrderDetails} from "./OrderDetails";

type CardProps = {
    order: IOrder | undefined
    reqResult: string
}

export const OrderCard: React.FC<CardProps> = (props) => {

    return (
        props.order
            ?
            <Card title={`Order with ${props.order.order_uid}`} style={{ width: '40%' }}>
                <Divider orientation="left">Order details</Divider>
                <OrderDetails details={props.order}/>
                <Divider orientation="left">Delivery details</Divider>
                <DeliveryDetails details={props.order.delivery}/>
                <Divider orientation="left">Payment details</Divider>
                <PaymentDetails details={props.order.payment}/>
                <Divider orientation="left">Items details</Divider>
                {props.order.items.map((item) => <ItemDetails details={item}/>)}
            </Card>
            :
            <p>
                {props.reqResult}
            </p>
    );
}