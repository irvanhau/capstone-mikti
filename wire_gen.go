// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"capstone-mikti/configs"
	"capstone-mikti/features/bookings"
	data7 "capstone-mikti/features/bookings/data"
	handler7 "capstone-mikti/features/bookings/handler"
	service7 "capstone-mikti/features/bookings/service"
	"capstone-mikti/features/categories"
	data3 "capstone-mikti/features/categories/data"
	handler3 "capstone-mikti/features/categories/handler"
	service3 "capstone-mikti/features/categories/service"
	"capstone-mikti/features/events"
	data2 "capstone-mikti/features/events/data"
	handler2 "capstone-mikti/features/events/handler"
	service2 "capstone-mikti/features/events/service"
	"capstone-mikti/features/payments"
	data8 "capstone-mikti/features/payments/data"
	handler8 "capstone-mikti/features/payments/handler"
	service8 "capstone-mikti/features/payments/service"
	"capstone-mikti/features/tickets"
	data5 "capstone-mikti/features/tickets/data"
	handler5 "capstone-mikti/features/tickets/handler"
	service5 "capstone-mikti/features/tickets/service"
	"capstone-mikti/features/users"
	"capstone-mikti/features/users/data"
	"capstone-mikti/features/users/handler"
	"capstone-mikti/features/users/service"
	"capstone-mikti/features/vouchers"
	data6 "capstone-mikti/features/vouchers/data"
	handler6 "capstone-mikti/features/vouchers/handler"
	service6 "capstone-mikti/features/vouchers/service"
	"capstone-mikti/features/wishlists"
	data4 "capstone-mikti/features/wishlists/data"
	handler4 "capstone-mikti/features/wishlists/handler"
	service4 "capstone-mikti/features/wishlists/service"
	"capstone-mikti/helper/email"
	"capstone-mikti/helper/enkrip"
	"capstone-mikti/helper/jwt"
	"capstone-mikti/routes"
	"capstone-mikti/server"
	"capstone-mikti/utils/cloudinary"
	"capstone-mikti/utils/database"
	"capstone-mikti/utils/midtrans"
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializedServer() *server.Server {
	programmingConfig := configs.InitConfig()
	db := database.InitDB(programmingConfig)
	hashInterface := enkrip.New()
	userData := data.New(db, hashInterface)
	jwtInterface := jwt.NewJWT(programmingConfig)
	emailInterface := email.New(programmingConfig)
	userService := service.New(userData, jwtInterface, hashInterface, emailInterface)
	userHandler := handler.NewHandler(userService, jwtInterface)
	eventData := data2.New(db)
	cloudinaryInterface := cloudinary.InitCloud(programmingConfig)
	categoryData := data3.New(db)
	eventService := service2.New(eventData, jwtInterface, cloudinaryInterface, categoryData)
	eventHandler := handler2.NewHandler(eventService, jwtInterface)
	categoryService := service3.New(categoryData)
	categoryHandler := handler3.NewHandler(categoryService, jwtInterface)
	wishlistData := data4.New(db)
	wishlistService := service4.New(wishlistData)
	wishlistHandler := handler4.NewHandler(wishlistService, jwtInterface)
	ticketData := data5.New(db)
	ticketService := service5.New(ticketData)
	ticketHandler := handler5.NewHandler(ticketService, jwtInterface)
	voucherData := data6.New(db)
	voucherService := service6.New(voucherData)
	voucherHandler := handler6.NewHandler(voucherService, jwtInterface)
	bookingData := data7.New(db)
	bookingService := service7.New(bookingData, jwtInterface)
	bookingHandler := handler7.NewHandler(bookingService, jwtInterface)
	paymentData := data8.New(db)
	midtransService := midtrans.InitMidtrans(programmingConfig)
	paymentService := service8.New(paymentData, midtransService, voucherData, ticketData)
	paymentHandler := handler8.NewHandler(paymentService, jwtInterface)
	echo := routes.NewRoute(programmingConfig, userHandler, eventHandler, categoryHandler, wishlistHandler, ticketHandler, voucherHandler, bookingHandler, paymentHandler)
	serverServer := server.InitServer(echo, programmingConfig)
	return serverServer
}

// injector.go:

var userSet = wire.NewSet(data.New, wire.Bind(new(users.UserDataInterface), new(*data.UserData)), service.New, wire.Bind(new(users.UserServiceInterface), new(*service.UserService)), handler.NewHandler, wire.Bind(new(users.UserHandlerInterface), new(*handler.UserHandler)))

var categorySet = wire.NewSet(data3.New, wire.Bind(new(categories.CategoryDataInterface), new(*data3.CategoryData)), service3.New, wire.Bind(new(categories.CategoryServiceInterface), new(*service3.CategoryService)), handler3.NewHandler, wire.Bind(new(categories.CategoryHandlerInterface), new(*handler3.CategoryHandler)))

var eventSet = wire.NewSet(data2.New, wire.Bind(new(events.EventDataInterface), new(*data2.EventData)), service2.New, wire.Bind(new(events.EventServiceInterface), new(*service2.EventService)), handler2.NewHandler, wire.Bind(new(events.EventHandlerInterface), new(*handler2.EventHandler)))

var voucherSet = wire.NewSet(data6.New, wire.Bind(new(vouchers.VoucherDataInterface), new(*data6.VoucherData)), service6.New, wire.Bind(new(vouchers.VoucherServiceInterface), new(*service6.VoucherService)), handler6.NewHandler, wire.Bind(new(vouchers.VoucherHandlerInterface), new(*handler6.VoucherHandler)))

var wishlistSet = wire.NewSet(data4.New, wire.Bind(new(wishlists.WishlistDataInterface), new(*data4.WishlistData)), service4.New, wire.Bind(new(wishlists.WishlistServiceInterface), new(*service4.WishlistService)), handler4.NewHandler, wire.Bind(new(wishlists.WishlistHandlerInterface), new(*handler4.WishlistHandler)))

var bookingSet = wire.NewSet(data7.New, wire.Bind(new(bookings.BookingDataInterface), new(*data7.BookingData)), service7.New, wire.Bind(new(bookings.BookingServiceInterface), new(*service7.BookingService)), handler7.NewHandler, wire.Bind(new(bookings.BookingHandlerInterface), new(*handler7.BookingHandler)))

var ticketSet = wire.NewSet(data5.New, wire.Bind(new(tickets.TicketDataInterface), new(*data5.TicketData)), service5.New, wire.Bind(new(tickets.TicketServiceInterface), new(*service5.TicketService)), handler5.NewHandler, wire.Bind(new(tickets.TicketHandlerInterface), new(*handler5.TicketHandler)))

var paymentSet = wire.NewSet(data8.New, wire.Bind(new(payments.PaymentDataInterface), new(*data8.PaymentData)), service8.New, wire.Bind(new(payments.PaymentServiceInterface), new(*service8.PaymentService)), handler8.NewHandler, wire.Bind(new(payments.PaymentHandlerInterface), new(*handler8.PaymentHandler)))
