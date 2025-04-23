package auth

import (
	"errors"
	"sync"
	"time"
)

const (
	// 缓存TTL，默认5分钟
	cacheTTL = 5 * time.Minute
)

// ApiKeyStruct API密钥结构
type ApiKeyStruct struct {
	AppId        string    `json:"api_key"`
	AppSecret    string    `json:"app_secret"`
	MerchantID   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	IsActive     bool      `json:"is_active"`
	ExpireAt     time.Time `json:"expire_at"`
	AllowedIPs   string    `json:"allowed_ips"`
	AllowedPaths string    `json:"allowed_paths"`
}

// ApiKeyManager API密钥管理器
type ApiKeyManager struct {
	keys      map[string]*ApiKeyStruct
	mutex     sync.RWMutex
	lastFetch time.Time
}

// NewApiKeyManager 创建新的API密钥管理器
func NewApiKeyManager() *ApiKeyManager {
	return &ApiKeyManager{
		keys:      make(map[string]*ApiKeyStruct),
		mutex:     sync.RWMutex{},
		lastFetch: time.Time{},
	}
}

// GetApiKey 获取API密钥信息
func (m *ApiKeyManager) GetApiKey(apiKey string) (*ApiKeyStruct, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if m.keys == nil {
		return nil, errors.New("API密钥缓存未初始化")
	}
	if key, ok := m.keys[apiKey]; ok {
		return key, nil
	}
	return nil, errors.New("API密钥不存在")
}

// RefreshCache 刷新缓存
func (m *ApiKeyManager) RefreshCache(keys map[string]*ApiKeyStruct) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.keys = keys
	m.lastFetch = time.Now()
}

// NeedRefresh 检查是否需要刷新缓存
func (m *ApiKeyManager) NeedRefresh() bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return time.Since(m.lastFetch) > cacheTTL
}
