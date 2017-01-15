package routes

import (
    "net/http"
    "domio_api/handlers/create_user_handler"
    "domio_api/handlers/login_user_handler"
    "domio_api/handlers/verify_token_handler"
    "domio_api/handlers/get_available_domains_handler"
    "domio_api/handlers/get_user_domains_handler"
    "domio_api/handlers/create_domain_handler"
    "domio_api/handlers/get_domain_info_handler"
    "domio_api/handlers/create_subscription_handler"
    "domio_api/handlers/create_card_handler"
    "domio_api/handlers/delete_domain_handler"
    "domio_api/handlers/delete_subscription_handler"
    "domio_api/handlers/get_user_subscriptions_handler"
    "domio_api/handlers/show_status_handler"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var RoutesList = Routes{
    Route{
        "ShowStatus",
        "GET",
        "/",
        show_status_handler.ShowStatusHandler,
    },
    Route{
        "CreateUser",
        "POST",
        "/users",
        create_user_handler.CreateUserHandler,
    },
    Route{
        "LoginUser",
        "POST",
        "/users/login",
        login_user_handler.LoginUserHandler,
    },
    Route{
        "VerifyToken",
        "POST",
        "/tokens/verify",
        verify_token_handler.VerifyTokenHandler,
    },
    Route{
        "GetAvailableDomains",
        "GET",
        "/domains/available",
        get_available_domains_handler.GetAvailableDomainsHandler,
    },
    Route{
        "GetUserDomains",
        "GET",
        "/domains/user",
        get_user_domains_handler.GetUserDomainsHandler,
    },
    Route{
        "CreateDomain",
        "POST",
        "/domains",
        create_domain_handler.CreateDomainHandler,
    },
    Route{
        "DeleteDomain",
        "DELETE",
        "/domains/{name}",
        delete_domain_handler.DeleteDomainHandler,
    },
    Route{
        "GetUserSubscriptions",
        "GET",
        "/subscriptions",
        get_user_subscriptions_handler.GetUserSubscriptionsHandler,
    },
    Route{
        "CreateSubscription",
        "POST",
        "/subscriptions",
        create_subscription_handler.CreateSubscriptionHandler,
    },
    Route{
        "DeleteSubscription",
        "DELETE",
        "/subscriptions/{id}",
        delete_subscription_handler.DeleteSubscriptionHandler,
    },
    Route{
        "CreateCard",
        "POST",
        "/cards",
        create_card_handler.CreateCardHandler,
    },
    Route{
        "GetDomainInfo",
        "GET",
        "/domains/{name}",
        get_domain_info_handler.GetDomainInfoHandler,
    },
}