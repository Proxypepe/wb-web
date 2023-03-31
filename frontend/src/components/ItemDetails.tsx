import {FC} from "react";
import {IItem} from "../models/Order";
import {Col, Row} from "antd";

type ItemDetailsProps = {
    details: IItem
}

export const ItemDetails: FC<ItemDetailsProps> = (props) => {

    return (
        <Row gutter={[8, 16]}>
            <Col span={12}>
                Chrt id
            </Col>
            <Col span={12}>
                {props.details.chrt_id}
            </Col>

            <Col span={12}>
                Track number
            </Col>
            <Col span={12}>
                {props.details.track_number}
            </Col>

            <Col span={12}>
                Price
            </Col>
            <Col span={12}>
                {props.details.price}
            </Col>

            <Col span={12}>
                Rid
            </Col>
            <Col span={12}>
                {props.details.rid}
            </Col>

            <Col span={12}>
                Name
            </Col>
            <Col span={12}>
                {props.details.name}
            </Col>

            <Col span={12}>
                Sale
            </Col>
            <Col span={12}>
                {props.details.sale}
            </Col>

            <Col span={12}>
                Size
            </Col>
            <Col span={12}>
                {props.details.size}
            </Col>

            <Col span={12}>
                Total price
            </Col>
            <Col span={12}>
                {props.details.total_price}
            </Col>

            <Col span={12}>
                nm id
            </Col>
            <Col span={12}>
                {props.details.nm_id}
            </Col>

            <Col span={12}>
                Brand
            </Col>
            <Col span={12}>
                {props.details.brand}
            </Col>

            <Col span={12}>
                Status
            </Col>
            <Col span={12}>
                {props.details.status}
            </Col>
        </Row>
    );
}