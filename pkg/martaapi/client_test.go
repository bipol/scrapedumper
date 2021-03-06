package martaapi_test

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"

	"github.com/smartatransit/scrapedumper/pkg/martaapi"
	. "github.com/smartatransit/scrapedumper/pkg/martaapi"
	"github.com/smartatransit/scrapedumper/pkg/martaapi/martaapifakes"
)

var _ = Describe("Client", func() {
	var (
		doer   *martaapifakes.FakeDoer
		apiKey string
		client Client
		resp   *http.Response
		retErr error
		err    error
	)
	BeforeEach(func() {
		resp = &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString("[]")),
		}
		doer = new(martaapifakes.FakeDoer)
		apiKey = "apikey"
		retErr = nil
		err = nil
	})
	JustBeforeEach(func() {
		doer.DoReturns(resp, retErr)
		logger, _ := zap.NewProduction()
		defer func() {
			_ = logger.Sync()
		}()
		client = New(
			doer,
			apiKey,
			logger,
			"test",
			"prefix",
		)
	})
	Context("New", func() {
		When("called", func() {
			It("does not return an err", func() {
				Expect(client).ToNot(BeNil())
			})
		})
	})
	Context("FindSchedules", func() {
		JustBeforeEach(func() {
			_, err = client.FindSchedules(context.Background())
		})
		When("the doer fails", func() {
			BeforeEach(func() {
				doer.DoReturns(nil, errors.New("do failed"))
			})
			It("fails", func() {
				Expect(err).To(HaveOccurred())
			})
		})
		When("the request receives a non-normal error code", func() {
			BeforeEach(func() {
				doer.DoReturns(&http.Response{StatusCode: http.StatusBadGateway}, nil)
			})
			It("fails", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

var _ = Describe("Schedule", func() {
	Describe("HasArrived", func() {
		It("works", func() {
			Expect(martaapi.Schedule{WaitingTime: "ArrIVIng"}.HasArrived()).To(BeFalse())
			Expect(martaapi.Schedule{WaitingTime: "arrIVED"}.HasArrived()).To(BeTrue())
			Expect(martaapi.Schedule{WaitingTime: "boarDING"}.HasArrived()).To(BeTrue())
			Expect(martaapi.Schedule{WaitingTime: "WHAT"}.HasArrived()).To(BeFalse())
		})
	})
	Describe("IsArriving", func() {
		It("works", func() {
			Expect(martaapi.Schedule{WaitingTime: "ArrIVIng"}.IsArriving()).To(BeTrue())
			Expect(martaapi.Schedule{WaitingTime: "arrIVED"}.IsArriving()).To(BeFalse())
			Expect(martaapi.Schedule{WaitingTime: "boarDING"}.IsArriving()).To(BeFalse())
			Expect(martaapi.Schedule{WaitingTime: "WHAT"}.IsArriving()).To(BeFalse())
		})
	})
	Describe("String", func() {
		It("works", func() {
			Expect(martaapi.Schedule{Direction: "N"}.String()).NotTo(BeEmpty())
		})
	})
})
