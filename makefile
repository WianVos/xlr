test_all: test_main test_data_model_ci test_data_model_template

test_main:
	go test -v
test_data_model_template:
	go test -v github.com/wianvos/xlr/datamodels/template

test_data_model_ci:
	go test -v github.com/wianvos/xlr/datamodels/ci
