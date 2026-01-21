-- Create Clients Table
CREATE TABLE public.clients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Fraud Rules Table
CREATE TABLE public.fraud_rules (
    id SERIAL PRIMARY KEY,
    client_id INTEGER NOT NULL,
    rule_name VARCHAR(100) NOT NULL,
    description TEXT,
    drl_content TEXT NOT NULL, -- The actual Grule DRL logic
    priority INTEGER DEFAULT 1, -- Execution order
    version INTEGER DEFAULT 1,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES public.clients (id) ON DELETE CASCADE
);

CREATE INDEX idx_fraud_rules_active ON public.fraud_rules (is_active);

CREATE TABLE public.data_templates (
    id SERIAL PRIMARY KEY,
    template_key VARCHAR(50) UNIQUE NOT NULL, -- e.g., 'monthly_biller_sum'
    query_sql TEXT NOT NULL, -- e.g., 'SELECT sum(amount) FROM tx WHERE user_id = :1'
    db_source VARCHAR(50) NOT NULL, -- Which DB to query (Biller_DB, Wallet_DB, etc.)
    cache_ttl INTERVAL DEFAULT '1 hour', -- How long to keep the result in Redis
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.transactions (
    id SERIAL PRIMARY KEY,
    client_id INTEGER NOT NULL,
    payload JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES public.clients (id) ON DELETE CASCADE
)

CREATE TABLE public.user_risk_profiles (
    user_id VARCHAR(100) PRIMARY KEY,
    overall_score DECIMAL(5, 2) DEFAULT 0, -- 0.00 to 100.00
    risk_level VARCHAR(20) DEFAULT 'LOW', -- LOW, MEDIUM, HIGH, CRITICAL
    features JSONB DEFAULT '{}', -- Persistent features like {"avg_tx": 500, "trust_level": 0.9}
    last_updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_user_risk_score ON public.user_risk_profiles (overall_score);

CREATE INDEX idx_user_risk_features ON public.user_risk_profiles USING GIN (features);

CREATE TABLE public.fraud_decisions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    transaction_id VARCHAR(100),
    user_id VARCHAR(100),
    decision VARCHAR(20), -- APPROVED, REJECTED, CHALLENGE
    rules_triggered JSONB, -- e.g., ["HighVelocity", "NewBillerCheck"]
    raw_features JSONB, -- The data state at the moment of decision
    latency_ms INTEGER, -- Performance tracking
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);