package routes

import (
    "net/http"
    "domio/handlers/create_user_handler"
    "domio/handlers/login_user_handler"
    "domio/handlers/verify_token_handler"
    "domio/handlers/get_available_domains_handler"
    "domio/handlers/get_user_domains_handler"
    "domio/handlers/create_domain_handler"
    "domio/handlers/get_domain_info_handler"
    "domio/handlers/create_subscription_handler"
    "domio/handlers/create_card_handler"
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
        "CreateUser",
        "POST",
        "/users",
        create_user_handler.CreateUserHandler,
    },
    Route{
        "LoginUser",
        "POST",
        "/users/login",
        login_user_handler.LoginUser,
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
        "CreateSubscription",
        "POST",
        "/subscriptions",
        create_subscription_handler.CreateSubscriptionHandler,
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