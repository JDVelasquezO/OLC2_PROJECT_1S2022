// Generated from /home/jdvelasquezo/GolandProjects/OLC2_Project1/server/ProjectLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ProjectLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PRINTLN=1, PRINT=2, DECLARAR=3, RSTRING=4, RINTEGER=5, RREAL=6, RBOOLEAN=7, 
		RSTR=8, RAS=9, RIF=10, RELSE=11, RAND=12, ROR=13, MULTILINE=14, INLINE=15, 
		INTEGER=16, FLOAT=17, CHAR=18, STRING=19, BOOLEAN=20, ID=21, EQUAL=22, 
		DOT=23, SEMICOLON=24, COMMA=25, COLON=26, ADMIRATION=27, REFERENCE=28, 
		LEFT_PAR=29, RIGHT_PAR=30, LEFT_KEY=31, RIGHT_KEY=32, LEFT_BRACKET=33, 
		RIGHT_BRACKET=34, EQUEAL_EQUAL=35, NOTEQUAL=36, GREATER_THAN=37, LESS_THAN=38, 
		GREATER_EQUALTHAN=39, LESS_EQUEALTHAN=40, MUL=41, DIV=42, ADD=43, SUB=44, 
		MOD=45, POWI=46, POWF=47, WHITESPACE=48;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PRINTLN", "PRINT", "DECLARAR", "RSTRING", "RINTEGER", "RREAL", "RBOOLEAN", 
			"RSTR", "RAS", "RIF", "RELSE", "RAND", "ROR", "MULTILINE", "INLINE", 
			"INTEGER", "FLOAT", "CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", 
			"SEMICOLON", "COMMA", "COLON", "ADMIRATION", "REFERENCE", "LEFT_PAR", 
			"RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET", 
			"EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN", 
			"LESS_EQUEALTHAN", "MUL", "DIV", "ADD", "SUB", "MOD", "POWI", "POWF", 
			"WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println'", "'print'", "'let'", "'String'", "'i64'", "'f64'", 
			"'Boolean'", "'str'", "'as'", "'If'", "'entonces'", "'AND'", "'OR'", 
			null, null, null, null, null, null, null, null, "'='", "'.'", "';'", 
			"','", "':'", "'!'", "'&'", "'('", "')'", "'{'", "'}'", "'['", "']'", 
			"'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'*'", "'/'", "'+'", "'-'", 
			"'%'", "'pow'", "'powf'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PRINTLN", "PRINT", "DECLARAR", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RSTR", "RAS", "RIF", "RELSE", "RAND", "ROR", "MULTILINE", 
			"INLINE", "INTEGER", "FLOAT", "CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", 
			"DOT", "SEMICOLON", "COMMA", "COLON", "ADMIRATION", "REFERENCE", "LEFT_PAR", 
			"RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET", 
			"EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN", 
			"LESS_EQUEALTHAN", "MUL", "DIV", "ADD", "SUB", "MOD", "POWI", "POWF", 
			"WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public ProjectLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ProjectLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\62\u0135\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\3\2\3\2\3\2\3"+
		"\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\13\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\17\3\17\3\17\3\17"+
		"\6\17\u00ad\n\17\r\17\16\17\u00ae\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3"+
		"\20\3\20\6\20\u00ba\n\20\r\20\16\20\u00bb\3\20\3\20\3\21\6\21\u00c1\n"+
		"\21\r\21\16\21\u00c2\3\22\6\22\u00c6\n\22\r\22\16\22\u00c7\3\22\3\22\6"+
		"\22\u00cc\n\22\r\22\16\22\u00cd\3\23\3\23\3\23\3\23\3\24\3\24\7\24\u00d6"+
		"\n\24\f\24\16\24\u00d9\13\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3"+
		"\25\3\25\3\25\5\25\u00e6\n\25\3\26\3\26\7\26\u00ea\n\26\f\26\16\26\u00ed"+
		"\13\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35"+
		"\3\35\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3$\3%\3%\3%"+
		"\3&\3&\3\'\3\'\3(\3(\3(\3)\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3"+
		"/\3/\3\60\3\60\3\60\3\60\3\60\3\61\6\61\u012d\n\61\r\61\16\61\u012e\3"+
		"\61\3\61\3\62\3\62\3\62\2\2\63\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61"+
		"\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61"+
		"a\62c\2\3\2\n\3\2\61\61\3\2\f\f\3\2\62;\3\2$$\7\2C\\aac|\u00d3\u00d3\u00f3"+
		"\u00f3\7\2\62;C\\c|\u00d3\u00d3\u00f3\u00f3\5\2\13\f\17\17\"\"\t\2\"#"+
		"%%--/\60<<BB]_\2\u013c\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2"+
		"\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25"+
		"\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2"+
		"\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2"+
		"\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3"+
		"\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2"+
		"\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2"+
		"Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3"+
		"\2\2\2\2_\3\2\2\2\2a\3\2\2\2\3e\3\2\2\2\5m\3\2\2\2\7s\3\2\2\2\tw\3\2\2"+
		"\2\13~\3\2\2\2\r\u0082\3\2\2\2\17\u0086\3\2\2\2\21\u008e\3\2\2\2\23\u0092"+
		"\3\2\2\2\25\u0095\3\2\2\2\27\u0098\3\2\2\2\31\u00a1\3\2\2\2\33\u00a5\3"+
		"\2\2\2\35\u00a8\3\2\2\2\37\u00b5\3\2\2\2!\u00c0\3\2\2\2#\u00c5\3\2\2\2"+
		"%\u00cf\3\2\2\2\'\u00d3\3\2\2\2)\u00e5\3\2\2\2+\u00e7\3\2\2\2-\u00ee\3"+
		"\2\2\2/\u00f0\3\2\2\2\61\u00f2\3\2\2\2\63\u00f4\3\2\2\2\65\u00f6\3\2\2"+
		"\2\67\u00f8\3\2\2\29\u00fa\3\2\2\2;\u00fc\3\2\2\2=\u00fe\3\2\2\2?\u0100"+
		"\3\2\2\2A\u0102\3\2\2\2C\u0104\3\2\2\2E\u0106\3\2\2\2G\u0108\3\2\2\2I"+
		"\u010b\3\2\2\2K\u010e\3\2\2\2M\u0110\3\2\2\2O\u0112\3\2\2\2Q\u0115\3\2"+
		"\2\2S\u0118\3\2\2\2U\u011a\3\2\2\2W\u011c\3\2\2\2Y\u011e\3\2\2\2[\u0120"+
		"\3\2\2\2]\u0122\3\2\2\2_\u0126\3\2\2\2a\u012c\3\2\2\2c\u0132\3\2\2\2e"+
		"f\7r\2\2fg\7t\2\2gh\7k\2\2hi\7p\2\2ij\7v\2\2jk\7n\2\2kl\7p\2\2l\4\3\2"+
		"\2\2mn\7r\2\2no\7t\2\2op\7k\2\2pq\7p\2\2qr\7v\2\2r\6\3\2\2\2st\7n\2\2"+
		"tu\7g\2\2uv\7v\2\2v\b\3\2\2\2wx\7U\2\2xy\7v\2\2yz\7t\2\2z{\7k\2\2{|\7"+
		"p\2\2|}\7i\2\2}\n\3\2\2\2~\177\7k\2\2\177\u0080\78\2\2\u0080\u0081\7\66"+
		"\2\2\u0081\f\3\2\2\2\u0082\u0083\7h\2\2\u0083\u0084\78\2\2\u0084\u0085"+
		"\7\66\2\2\u0085\16\3\2\2\2\u0086\u0087\7D\2\2\u0087\u0088\7q\2\2\u0088"+
		"\u0089\7q\2\2\u0089\u008a\7n\2\2\u008a\u008b\7g\2\2\u008b\u008c\7c\2\2"+
		"\u008c\u008d\7p\2\2\u008d\20\3\2\2\2\u008e\u008f\7u\2\2\u008f\u0090\7"+
		"v\2\2\u0090\u0091\7t\2\2\u0091\22\3\2\2\2\u0092\u0093\7c\2\2\u0093\u0094"+
		"\7u\2\2\u0094\24\3\2\2\2\u0095\u0096\7K\2\2\u0096\u0097\7h\2\2\u0097\26"+
		"\3\2\2\2\u0098\u0099\7g\2\2\u0099\u009a\7p\2\2\u009a\u009b\7v\2\2\u009b"+
		"\u009c\7q\2\2\u009c\u009d\7p\2\2\u009d\u009e\7e\2\2\u009e\u009f\7g\2\2"+
		"\u009f\u00a0\7u\2\2\u00a0\30\3\2\2\2\u00a1\u00a2\7C\2\2\u00a2\u00a3\7"+
		"P\2\2\u00a3\u00a4\7F\2\2\u00a4\32\3\2\2\2\u00a5\u00a6\7Q\2\2\u00a6\u00a7"+
		"\7T\2\2\u00a7\34\3\2\2\2\u00a8\u00a9\7\61\2\2\u00a9\u00aa\7,\2\2\u00aa"+
		"\u00ac\3\2\2\2\u00ab\u00ad\n\2\2\2\u00ac\u00ab\3\2\2\2\u00ad\u00ae\3\2"+
		"\2\2\u00ae\u00ac\3\2\2\2\u00ae\u00af\3\2\2\2\u00af\u00b0\3\2\2\2\u00b0"+
		"\u00b1\7,\2\2\u00b1\u00b2\7\61\2\2\u00b2\u00b3\3\2\2\2\u00b3\u00b4\b\17"+
		"\2\2\u00b4\36\3\2\2\2\u00b5\u00b6\7\61\2\2\u00b6\u00b7\7\61\2\2\u00b7"+
		"\u00b9\3\2\2\2\u00b8\u00ba\n\3\2\2\u00b9\u00b8\3\2\2\2\u00ba\u00bb\3\2"+
		"\2\2\u00bb\u00b9\3\2\2\2\u00bb\u00bc\3\2\2\2\u00bc\u00bd\3\2\2\2\u00bd"+
		"\u00be\b\20\2\2\u00be \3\2\2\2\u00bf\u00c1\t\4\2\2\u00c0\u00bf\3\2\2\2"+
		"\u00c1\u00c2\3\2\2\2\u00c2\u00c0\3\2\2\2\u00c2\u00c3\3\2\2\2\u00c3\"\3"+
		"\2\2\2\u00c4\u00c6\t\4\2\2\u00c5\u00c4\3\2\2\2\u00c6\u00c7\3\2\2\2\u00c7"+
		"\u00c5\3\2\2\2\u00c7\u00c8\3\2\2\2\u00c8\u00c9\3\2\2\2\u00c9\u00cb\7\60"+
		"\2\2\u00ca\u00cc\t\4\2\2\u00cb\u00ca\3\2\2\2\u00cc\u00cd\3\2\2\2\u00cd"+
		"\u00cb\3\2\2\2\u00cd\u00ce\3\2\2\2\u00ce$\3\2\2\2\u00cf\u00d0\7)\2\2\u00d0"+
		"\u00d1\n\5\2\2\u00d1\u00d2\7)\2\2\u00d2&\3\2\2\2\u00d3\u00d7\7$\2\2\u00d4"+
		"\u00d6\n\5\2\2\u00d5\u00d4\3\2\2\2\u00d6\u00d9\3\2\2\2\u00d7\u00d5\3\2"+
		"\2\2\u00d7\u00d8\3\2\2\2\u00d8\u00da\3\2\2\2\u00d9\u00d7\3\2\2\2\u00da"+
		"\u00db\7$\2\2\u00db(\3\2\2\2\u00dc\u00dd\7v\2\2\u00dd\u00de\7t\2\2\u00de"+
		"\u00df\7w\2\2\u00df\u00e6\7g\2\2\u00e0\u00e1\7h\2\2\u00e1\u00e2\7c\2\2"+
		"\u00e2\u00e3\7n\2\2\u00e3\u00e4\7u\2\2\u00e4\u00e6\7g\2\2\u00e5\u00dc"+
		"\3\2\2\2\u00e5\u00e0\3\2\2\2\u00e6*\3\2\2\2\u00e7\u00eb\t\6\2\2\u00e8"+
		"\u00ea\t\7\2\2\u00e9\u00e8\3\2\2\2\u00ea\u00ed\3\2\2\2\u00eb\u00e9\3\2"+
		"\2\2\u00eb\u00ec\3\2\2\2\u00ec,\3\2\2\2\u00ed\u00eb\3\2\2\2\u00ee\u00ef"+
		"\7?\2\2\u00ef.\3\2\2\2\u00f0\u00f1\7\60\2\2\u00f1\60\3\2\2\2\u00f2\u00f3"+
		"\7=\2\2\u00f3\62\3\2\2\2\u00f4\u00f5\7.\2\2\u00f5\64\3\2\2\2\u00f6\u00f7"+
		"\7<\2\2\u00f7\66\3\2\2\2\u00f8\u00f9\7#\2\2\u00f98\3\2\2\2\u00fa\u00fb"+
		"\7(\2\2\u00fb:\3\2\2\2\u00fc\u00fd\7*\2\2\u00fd<\3\2\2\2\u00fe\u00ff\7"+
		"+\2\2\u00ff>\3\2\2\2\u0100\u0101\7}\2\2\u0101@\3\2\2\2\u0102\u0103\7\177"+
		"\2\2\u0103B\3\2\2\2\u0104\u0105\7]\2\2\u0105D\3\2\2\2\u0106\u0107\7_\2"+
		"\2\u0107F\3\2\2\2\u0108\u0109\7?\2\2\u0109\u010a\7?\2\2\u010aH\3\2\2\2"+
		"\u010b\u010c\7#\2\2\u010c\u010d\7?\2\2\u010dJ\3\2\2\2\u010e\u010f\7@\2"+
		"\2\u010fL\3\2\2\2\u0110\u0111\7>\2\2\u0111N\3\2\2\2\u0112\u0113\7@\2\2"+
		"\u0113\u0114\7?\2\2\u0114P\3\2\2\2\u0115\u0116\7>\2\2\u0116\u0117\7?\2"+
		"\2\u0117R\3\2\2\2\u0118\u0119\7,\2\2\u0119T\3\2\2\2\u011a\u011b\7\61\2"+
		"\2\u011bV\3\2\2\2\u011c\u011d\7-\2\2\u011dX\3\2\2\2\u011e\u011f\7/\2\2"+
		"\u011fZ\3\2\2\2\u0120\u0121\7\'\2\2\u0121\\\3\2\2\2\u0122\u0123\7r\2\2"+
		"\u0123\u0124\7q\2\2\u0124\u0125\7y\2\2\u0125^\3\2\2\2\u0126\u0127\7r\2"+
		"\2\u0127\u0128\7q\2\2\u0128\u0129\7y\2\2\u0129\u012a\7h\2\2\u012a`\3\2"+
		"\2\2\u012b\u012d\t\b\2\2\u012c\u012b\3\2\2\2\u012d\u012e\3\2\2\2\u012e"+
		"\u012c\3\2\2\2\u012e\u012f\3\2\2\2\u012f\u0130\3\2\2\2\u0130\u0131\b\61"+
		"\2\2\u0131b\3\2\2\2\u0132\u0133\7^\2\2\u0133\u0134\t\t\2\2\u0134d\3\2"+
		"\2\2\f\2\u00ae\u00bb\u00c2\u00c7\u00cd\u00d7\u00e5\u00eb\u012e\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}