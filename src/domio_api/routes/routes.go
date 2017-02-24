package routes

import (
    "domio_api/handlers/create_user_handler"
    "domio_api/handlers/login_user_handler"
    "domio_api/types"
    "domio_api/handlers/verify_token_handler"
    "domio_api/handlers/get_available_domains_handler"
    "domio_api/handlers/get_user_domains_handler"
    "domio_api/handlers/get_user_card_handler"
    "domio_api/handlers/delete_card_handler"
    "domio_api/handlers/create_domain_handler"
    "domio_api/handlers/update_domain_handler"
    "domio_api/handlers/delete_domain_handler"
    "domio_api/handlers/get_user_subscriptions_handler"
    "domio_api/handlers/create_subscription_handler"
    "domio_api/handlers/update_subscription_handler"
    "domio_api/handlers/get_subscription_handler"
    "domio_api/handlers/delete_subscription_handler"
    "domio_api/handlers/delete_record_handler"
    "domio_api/handlers/create_card_handler"
    "domio_api/handlers/get_subscription_records_handler"
    "domio_api/handlers/update_subscription_records_handler"
    "domio_api/handlers/get_domain_info_handler"
    "domio_api/handlers/get_user_cards_handler"
)

type Routes []*types.Route

var RoutesList = Routes{
    create_user_handler.GetRoute(),
    login_user_handler.GetRoute(),
    get_domain_info_handler.GetRoute(),
    verify_token_handler.GetRoute(),
    get_available_domains_handler.GetRoute(),
    get_user_domains_handler.GetRoute(),
    get_user_cards_handler.GetRoute(),
    get_user_card_handler.GetRoute(),
    delete_card_handler.GetRoute(),
    create_domain_handler.GetRoute(),
    update_domain_handler.GetRoute(),
    delete_domain_handler.GetRoute(),
    get_user_subscriptions_handler.GetRoute(),
    create_subscription_handler.GetRoute(),
    update_subscription_handler.GetRoute(),
    get_subscription_handler.GetRoute(),
    delete_subscription_handler.GetRoute(),
    delete_record_handler.GetRoute(),
    create_card_handler.GetRoute(),
    get_subscription_records_handler.GetRoute(),
    update_subscription_records_handler.GetRoute(),
}