package models

import (
	"errors"
	"fmt"
	"github.com/otc/otc-web/utils/db_util"
	"gorm.io/gorm"
)

type Config struct {
	Id          int    `json:"id" gorm:"column:id"`
	ConfigKey   string `json:"config_key" gorm:"column:config_key"`
	ConfigValue string `json:"config_value" gorm:"column:config_value"`
	Comment     string `json:"comment" gorm:"column:comment"`
}

func (e *Config) TableName() string {
	return "config"
}

func (e *Config) GetEthUsdtCollectThreshold() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_USDT_COLLECT_THRESHOLD").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_USDT_COLLECT_THRESHOLD配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthUsdtCollectAddress() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_USDT_COLLECT_ADDRESS").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_USDT_COLLECT_ADDRESS配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthGasPrivateKey() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_GAS_PRIVATE_KEY").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_GAS_PRIVATE_KEY配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthUsdtWithdrawPrivateKey() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_USDT_WITHDRAW_PRIVATE_KEY").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_USDT_WITHDRAW_PRIVATE_KEY配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthUsdtContractAddress() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_USDT_CONTRACT_ADDRESS").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_USDT_CONTRACT_ADDRESS配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthGasAddress() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_GAS_ADDRESS").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_GAS_ADDRESS配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthUsdtWithdrawAddress() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_USDT_WITHDRAW_ADDRESS").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_USDT_WITHDRAW_ADDRESS配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthUsdtCollectFlag() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_USDT_COLLECT_FLAG").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_USDT_COLLECT_FLAG配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthUsdtCollectStatus() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_USDT_COLLECT_STATUS").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_USDT_COLLECT_STATUS配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthUsdtCollectStartTime() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_USDT_COLLECT_START_TIME").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_USDT_COLLECT_START_TIME配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthUsdtCollectInterval() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_USDT_COLLECT_INTERVAL").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_USDT_COLLECT_PERIOD配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthClientUrl() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_CLIENT_URL").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_CLIENT_URL配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthPrice() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_PRICE").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_CLIENT_URL配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetEthUsdtTimeGapBetweenGasAndCollect() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "ETH_USDT_TIME_GAP_BETWEEN_GAS_AND_COLLECT").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到ETH_CLIENT_URL配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

//TRX
func (e *Config) GetTrxUsdtCollectThreshold() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_USDT_COLLECT_THRESHOLD").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_USDT_COLLECT_THRESHOLD配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxUsdtCollectAddress() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_USDT_COLLECT_ADDRESS").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_USDT_COLLECT_ADDRESS配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxGasPrivateKey() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_GAS_PRIVATE_KEY").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_GAS_PRIVATE_KEY配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxUsdtWithdrawPrivateKey() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_USDT_WITHDRAW_PRIVATE_KEY").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_USDT_WITHDRAW_PRIVATE_KEY配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxUsdtContractAddress() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_USDT_CONTRACT_ADDRESS").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_USDT_CONTRACT_ADDRESS配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxGasAddress() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_GAS_ADDRESS").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_GAS_ADDRESS配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxUsdtWithdrawAddress() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_USDT_WITHDRAW_ADDRESS").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_USDT_WITHDRAW_ADDRESS配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxUsdtCollectFlag() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_USDT_COLLECT_FLAG").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_USDT_COLLECT_FLAG配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxUsdtCollectStatus() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_USDT_COLLECT_STATUS").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_USDT_COLLECT_STATUS配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxUsdtCollectStartTime() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_USDT_COLLECT_START_TIME").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_USDT_COLLECT_START_TIME配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxUsdtCollectInterval() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_USDT_COLLECT_INTERVAL").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_USDT_COLLECT_PERIOD配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxClientUrl() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_CLIENT_URL").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_CLIENT_URL配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxPrice() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_PRICE").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_CLIENT_URL配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}

func (e *Config) GetTrxUsdtTimeGapBetweenGasAndCollect() (string, error) {
	var config Config
	db := db_util.GetDB()
	tx := db.Debug().Model(&Config{}).Where("config_key = ?", "TRX_USDT_TIME_GAP_BETWEEN_GAS_AND_COLLECT").First(&config)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没找到TRX_CLIENT_URL配置")
		return "", tx.Error
	} else {
		return config.ConfigValue, nil
	}
}
